package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type cell struct {
	drawable uint32

	alive     bool
	aliveNext bool

	x int
	y int
}

const (
	rows               = 100
	columns            = 100
	width              = 500
	height             = 500
	fps                = 5
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 1);
    }
` + "\x00"
	threshold = 0.15
)

var (
	triangle = []float32{
		-0.5, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()

	cells := makeCells()

	for !window.ShouldClose() {
		t := time.Now()

		for x := range cells {
			for _, c := range cells[x] {
				c.checkState(cells)
			}
		}

		draw(cells, window, program)
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func initGlfw() *glfw.Window {

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("opengl version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func draw(cells [][]*cell, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for x := range cells {
		for _, c := range cells[x] {
			c.draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLen)

		log := strings.Repeat("\x00", int(logLen+1))
		gl.GetShaderInfoLog(shader, logLen, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile shader %v: %v", source, log)
	}

	return shader, nil
}

func makeCells() [][]*cell {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	cells := make([][]*cell, rows, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := newCell(x, y)
			c.alive = rand.Float64() < threshold
			c.aliveNext = c.alive
			cells[x] = append(cells[x], c)
		}
	}
	return cells
}

func newCell(x, y int) *cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var p float32
		var s float32

		switch i % 3 {
		case 0:
			s = 1.0 / float32(columns)
			p = float32(x) * s
		case 1:
			s = 1.0 / float32(rows)
			p = float32(y) * s
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (p * 2) - 1
		} else {
			points[i] = ((p + s) * 2) - 1
		}
	}

	return &cell{
		drawable: makeVao(points),
		x:        x,
		y:        y,
	}
}

func (c *cell) draw() {

	if !c.alive {
		return
	}

	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

// checkState determines the state of the cell for the next tick of the game.
func (c *cell) checkState(cells [][]*cell) {
	c.alive = c.aliveNext
	c.aliveNext = c.alive

	liveCount := c.liveNeighbors(cells)
	if c.alive {
		// 1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
		if liveCount < 2 {
			c.aliveNext = false
		}

		// 2. Any live cell with two or three live neighbours lives on to the next generation.
		if liveCount == 2 || liveCount == 3 {
			c.aliveNext = true
		}

		// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
		if liveCount > 3 {
			c.aliveNext = false
		}
	} else {
		// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		if liveCount == 3 {
			c.aliveNext = true
		}
	}
}

// liveNeighbors returns the number of live neighbors for a cell.
func (c *cell) liveNeighbors(cells [][]*cell) int {
	var liveCount int
	add := func(x, y int) {
		// If we're at an edge, check the other side of the board.
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].alive {
			liveCount++
		}
	}

	add(c.x-1, c.y)   // To the left
	add(c.x+1, c.y)   // To the right
	add(c.x, c.y+1)   // up
	add(c.x, c.y-1)   // down
	add(c.x-1, c.y+1) // top-left
	add(c.x+1, c.y+1) // top-right
	add(c.x-1, c.y-1) // bottom-left
	add(c.x+1, c.y-1) // bottom-right

	return liveCount
}
