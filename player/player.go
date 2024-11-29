package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 640
)

type Player struct {
	Position, Size rl.Vector2
	Life           int
}

func (p *Player) Spawn() {
	p.Position = rl.Vector2{float32(ScreenWidth / 2), float32(ScreenHeight * 7 / 8)}
	p.Size = rl.Vector2{float32(ScreenWidth / 10), 20}
}

func (p *Player) Move() {

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= 5
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.Position.X += 5
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.Position.Y -= 5
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += 5
	}
}
