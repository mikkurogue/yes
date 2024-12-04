package player

import (
	"fmt"
	"gengine/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	Attached bool
	Instance rl.Camera2D
}

func (c *Camera) Init() {

	if !c.Attached {
		fmt.Printf("WARNING:: no attached camera instance found before render draw")
		return
	}

	rl.BeginMode2D(c.Instance)
}

func (c *Camera) AttachNewCamera(player Player) {
	c.Attached = true

	c.Instance = rl.Camera2D{
		Offset:   rl.NewVector2(float32(constants.ScreenWidth/2), constants.ScreenHeight/2),
		Target:   player.Position,
		Rotation: 0.0,
		Zoom:     1.0,
	}
}

func (c *Camera) SetTarget(player Player) {
	c.Instance.Target = player.Position
}
