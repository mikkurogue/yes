package main

import (
	"gengine/constants"
	"gengine/enemy"
	"gengine/engine/collision"
	"gengine/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	pause       bool
	player      player.Player
	enemies     []enemy.Enemy
	projectiles []player.Projectile
	camera      player.Camera
}

var staticObjects = []collision.RigidBody2D{
	{
		Position: rl.Vector2{
			X: 300,
			Y: 300,
		},
		Size: rl.Vector2{
			X: 50,
			Y: 50,
		},
		IsStatic: true,
	},
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.player.Spawn()
	e := enemy.Enemy{}
	e.Spawn()
	g.enemies = append(g.enemies, e)

	g.camera.AttachNewCamera(g.player)
}

func (g *Game) Update() {
	deltaTime := rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeyP) {
		g.pause = !g.pause
	}

	if g.pause {
		return
	}

	g.player.Update(staticObjects, deltaTime)
	g.camera.SetTarget(g.player)

	if rl.IsKeyPressed(rl.KeySpace) {
		projectile := player.Projectile{}
		projectile.Init(g.player.Position, rl.Vector2{X: 0, Y: -1}) // Default direction: upward
		g.projectiles = append(g.projectiles, projectile)
	}

	for i := 0; i < len(g.projectiles); {
		g.projectiles[i].Move(deltaTime)
		projectileRemoved := false

		for j := 0; j < len(g.enemies); j++ {
			if g.projectiles[i].CheckCollision(&g.enemies[j]) {
				g.enemies[j].TakeDamage(50)
				g.projectiles[i].Spawned = false
				projectileRemoved = true // flag
				break                    // Stop checking further enemies for this projectile
			}
		}

		if projectileRemoved || !g.projectiles[i].Spawned {
			g.projectiles = append(g.projectiles[:i], g.projectiles[i+1:]...)
		} else {
			i++
		}
	}

	// Update enemies
	for i := 0; i < len(g.enemies); {
		if g.enemies[i].Health <= 0 {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		} else {
			i++
		}
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	if !g.pause {
		g.camera.Init()
		g.player.Draw()

		for _, obj := range staticObjects {
			DebugDreawCollisionBox(obj, rl.Red)
			rl.DrawRectangle(
				int32(obj.Position.X),
				int32(obj.Position.Y),
				int32(obj.Size.X),
				int32(obj.Size.Y),
				rl.Blue,
			)
		}

		for _, e := range g.enemies {
			e.Draw()
		}

		for _, p := range g.projectiles {
			if p.Spawned {
				p.Draw()
			}
		}

		rl.EndMode2D()
	}

	if g.pause {
		rl.DrawText("GAME PAUSED", constants.ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, constants.ScreenHeight/2+constants.ScreenHeight/4-40, 40, rl.Gray)
	}

	rl.EndDrawing()
}

func DebugDreawCollisionBox(rb collision.RigidBody2D, color rl.Color) {
	rl.DrawRectangleLines(
		int32(rb.Position.X),
		int32(rb.Position.Y),
		int32(rb.Size.X),
		int32(rb.Size.Y),
		color,
	)
}
