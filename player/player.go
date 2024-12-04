package player

import (
	"gengine/constants"
	"gengine/engine/collision"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	collision.RigidBody2D
	Health int
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
	p.Velocity = rl.Vector2{}
	p.IsStatic = false
}

func (p *Player) Update(allObjects []collision.RigidBody2D, dt float32) {
	p.Move()
	p.UpdatePhysics(allObjects, dt)
}

func (p *Player) Draw() {
	// Draw player
	rl.DrawRectangle(
		int32(p.Position.X-p.Size.X/2),
		int32(p.Position.Y-p.Size.Y/2),
		int32(p.Size.X),
		int32(p.Size.X),
		rl.DarkPurple,
	)

}

func (p *Player) Move() {

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.Velocity.X -= 5
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.Velocity.X += 5
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.Velocity.Y -= 5
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.Velocity.Y += 5
	}
}

func (p *Player) UpdatePhysics(allObjects []collision.RigidBody2D, dt float32) {
	p.Position.X += p.Velocity.X * dt
	p.Position.Y += p.Velocity.Y * dt

	for _, obj := range allObjects {
		if &p.RigidBody2D == &obj || obj.IsStatic == false {
			continue
		}

		if collision.CheckCollision(p.RigidBody2D, obj) {
			collision.ResolveCollision(&p.RigidBody2D, &obj)
		}
	}
}
