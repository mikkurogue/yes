package player

import (
	"gengine/enemy"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Speed               float32
	Position, Direction rl.Vector2
	Spawned             bool
	Lifetime            float32
	CollisionRadius     float32
}

func (p *Projectile) Init(position, direction rl.Vector2) {
	p.Position = position
	p.Speed = 15.0 // Set a default speed
	p.Direction = rl.Vector2Normalize(direction)
	p.Spawned = true
	p.Lifetime = 2.0        // spawn for 2 seconds
	p.CollisionRadius = 5.0 // radius as this is a circle
}

func (p *Projectile) Draw() {
	rl.DrawCircleV(p.Position, 5, rl.Red)
}

// Move the projectile in the p.Direction with the speed of p.Speed
func (p *Projectile) Move(deltaTime float32) {
	if p.Spawned {
		p.Position.X += p.Direction.X * p.Speed
		p.Position.Y += p.Direction.Y * p.Speed

		p.Lifetime -= deltaTime
		if p.Lifetime <= 0 {
			p.Spawned = false
		}
	}
}

func (p *Projectile) CheckCollision(enemy *enemy.Enemy) bool {
	if !p.Spawned {
		return false
	}

	return rl.CheckCollisionCircles(
		p.Position,
		p.CollisionRadius,
		enemy.Position,
		25, // enemy radius matching DrawCircleV
	)
}
