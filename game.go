package main

import (
	"gengine/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 640
)

type Game struct {
	pause  bool
	player player.Player
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.player.Position = rl.Vector2{float32(ScreenWidth / 2), float32(ScreenHeight * 7 / 8)}
	g.player.Size = rl.Vector2{float32(ScreenWidth / 10), 20}
}

func (g *Game) Update() {
	if !g.pause {
		g.player.Move()
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	if !g.pause {
		rl.DrawRectangle(
			int32(g.player.Position.X-g.player.Size.X/2),
			int32(g.player.Position.Y-g.player.Size.Y/2),
			int32(g.player.Size.X),
			int32(g.player.Size.Y), rl.Black,
		)
	}

	if g.pause {
		rl.DrawText("GAME PAUSED", ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, ScreenHeight/2+ScreenHeight/4-40, 40, rl.Gray)
	}

	rl.EndDrawing()
}
