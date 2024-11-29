package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position, Size rl.Vector2
	Life           int
}

func Spawn() {}

func (p *Player) Move() {

	// implement pause logic somehow

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
