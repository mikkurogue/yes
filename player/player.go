package player

import (
	"gengine/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position, Size rl.Vector2
	Health         int
}

func (p *Player) Spawn() {

	p.Health = 100.0

	p.Position = rl.Vector2{
		X: float32(constants.ScreenWidth / 2),
		Y: float32(constants.ScreenHeight * 7 / 8),
	}
	p.Size = rl.Vector2{
		X: float32(constants.ScreenWidth / 10),
		Y: 20.0,
	}
}

func (p *Player) Update() {
	p.Move()
}

func (p *Player) Draw() {
	// Draw player
	rl.DrawRectangle(
		int32(p.Position.X-p.Size.X/2),
		int32(p.Position.Y-p.Size.Y/2),
		int32(p.Size.X),
		int32(p.Size.X), rl.DarkPurple,
	)

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
