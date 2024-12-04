package main

import (
	"gengine/constants"
	"gengine/enemy"
	"gengine/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	pause       bool
	player      player.Player
	enemies     []enemy.Enemy
	projectiles []player.Projectile
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

	// for i := 0; i < 5; i++ { // Example: Spawn 5 enemies
	// 	newEnemy := enemy.Enemy{}
	// 	newEnemy.Spawn()
	// 	g.enemies = append(g.enemies, newEnemy)
	// }
}

func (g *Game) Update() {
	deltaTime := rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeyP) {
		g.pause = !g.pause
	}

	if g.pause {
		return
	}

	g.player.Update()

	if rl.IsKeyPressed(rl.KeySpace) {
		projectile := player.Projectile{}
		projectile.Init(g.player.Position, rl.Vector2{X: 0, Y: -1}) // Default direction: upward
		g.projectiles = append(g.projectiles, projectile)
	}

	for i := 0; i < len(g.projectiles); {
		g.projectiles[i].Move(deltaTime)
		projectileRemoved := false

		for _, e := range g.enemies {
			if g.projectiles[i].CheckCollision(&e) {
				e.TakeDamage(50)
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
		g.player.Draw()

		for _, e := range g.enemies {
			e.Draw()
		}

		for _, p := range g.projectiles {
			if p.Spawned {
				p.Draw()
			}
		}
	}

	if g.pause {
		rl.DrawText("GAME PAUSED", constants.ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, constants.ScreenHeight/2+constants.ScreenHeight/4-40, 40, rl.Gray)
	}

	rl.EndDrawing()
}
