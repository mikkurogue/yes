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
	p.Health = 100

	p.Position = rl.Vector2{
		X: float32(constants.ScreenWidth / 2),
		Y: float32(constants.ScreenHeight * 7 / 8),
	}

	// Set the player size to a square
	sideLength := float32(constants.ScreenWidth / 10) // Example: Adjust as needed
	p.Size = rl.Vector2{
		X: sideLength,
		Y: sideLength, // Make it a square
	}

	p.Velocity = rl.Vector2{}
	p.IsStatic = false
}

func (p *Player) Update(allObjects []collision.RigidBody2D, dt float32) {
	p.Move()
	p.UpdatePhysics(allObjects, dt)
}

func (p *Player) Draw() {
	// Draw player (centered rectangle)
	rl.DrawRectangle(
		int32(p.Position.X-p.Size.X/2),
		int32(p.Position.Y-p.Size.Y/2),
		int32(p.Size.X),
		int32(p.Size.Y), // Use p.Size.Y for height
		rl.DarkPurple,
	)

	// Draw collision box (for debugging)
	hitbox := collision.RigidBody2D{
		Position: rl.Vector2{
			X: p.Position.X - p.Size.X/2,
			Y: p.Position.Y - p.Size.Y/2,
		},
		Size: p.Size,
	}
	rl.DrawRectangleLines(
		int32(hitbox.Position.X),
		int32(hitbox.Position.Y),
		int32(hitbox.Size.X),
		int32(hitbox.Size.Y),
		rl.Green,
	)
}

func (p *Player) Move() {
	const speed float32 = 500.0 // px per sec

	p.Velocity = rl.Vector2{} // make sure to reset velocity when adding new inp command

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.Velocity.X = -speed
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.Velocity.X = speed
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.Velocity.Y = -speed
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.Velocity.Y = speed
	}
}

func (p *Player) UpdatePhysics(allObjects []collision.RigidBody2D, dt float32) {
	// Calculate the intended new position
	newPosition := rl.Vector2{
		X: p.Position.X + p.Velocity.X*dt,
		Y: p.Position.Y + p.Velocity.Y*dt,
	}

	// Check for collisions
	for _, obj := range allObjects {
		if &p.RigidBody2D == &obj || !obj.IsStatic {
			continue // Skip self and non-static objects
		}

		// Horizontal collision
		testBoxX := collision.RigidBody2D{
			Position: rl.Vector2{
				X: newPosition.X - p.Size.X/2,
				Y: p.Position.Y - p.Size.Y/2,
			},
			Size: p.Size,
		}
		if collision.CheckCollision(testBoxX, obj) {
			if p.Velocity.X > 0 { // Moving right
				newPosition.X = obj.Position.X - obj.Size.X/2 - p.Size.X/2
			} else if p.Velocity.X < 0 { // Moving left
				newPosition.X = obj.Position.X + obj.Size.X/2 + p.Size.X/2
			}
			p.Velocity.X = 0 // Stop horizontal movement
		}

		// Vertical collision
		testBoxY := collision.RigidBody2D{
			Position: rl.Vector2{
				X: p.Position.X - p.Size.X/2,
				Y: newPosition.Y - p.Size.Y/2,
			},
			Size: p.Size,
		}
		if collision.CheckCollision(testBoxY, obj) {
			if p.Velocity.Y > 0 { // Moving down
				newPosition.Y = obj.Position.Y - obj.Size.Y/2 - p.Size.Y/2
			} else if p.Velocity.Y < 0 { // Moving up
				newPosition.Y = obj.Position.Y + obj.Size.Y/2 + p.Size.Y/2
			}
			p.Velocity.Y = 0 // Stop vertical movement
		}
	}

	// Update player position to resolved position
	p.Position = newPosition
}
