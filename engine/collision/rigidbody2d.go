package collision

import rl "github.com/gen2brain/raylib-go/raylib"

type RigidBody2D struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Size     rl.Vector2
	IsStatic bool // determines if object is immovable (idk yet wtf i mean)
}

func CheckCollision(a, b RigidBody2D) bool {
	return a.Position.X < b.Position.X+b.Size.X &&
		a.Position.X+a.Size.X > b.Position.X &&
		a.Position.Y < b.Position.Y+b.Size.Y &&
		a.Position.Y+a.Size.Y > b.Position.Y
}

func ResolveCollision(a, b *RigidBody2D) {
	if !b.IsStatic {
		return // only handle collisions against statics frnow
	}

	// undo last movement to resolve collision
	if a.Velocity.X > 0 { // moving R
		a.Position.X = b.Position.X - a.Size.X
	} else if a.Velocity.X < 0 { // moving L
		a.Position.X = b.Position.X + b.Size.X
	}

	if a.Velocity.Y > 0 { // moving D
		a.Position.Y = b.Position.Y - a.Size.Y
	} else if a.Velocity.Y < 0 { // moving U
		a.Position.Y = b.Position.Y + b.Size.Y
	}

	a.Velocity = rl.Vector2{X: 0.0, Y: 0.0}
}

func UpdateRigidBody(rb *RigidBody2D, dt float32, allObjects []RigidBody2D) {

	if rb.IsStatic {
		// static objs dont move
		return
	}

	rb.Position.X += rb.Velocity.X * dt
	rb.Position.Y += rb.Velocity.Y * dt

	for _, other := range allObjects {
		if rb == &other {
			continue
		}
		if CheckCollision(*rb, other) {
			ResolveCollision(rb, &other)
		}
	}

}
