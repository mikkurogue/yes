package main

import (
	"gengine/constants"
	"gengine/player"

	rl "github.com/gen2brain/raylib-go/raylib"
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
	g.player.Spawn()
}

var projectiles []player.Projectile

func (g *Game) Update() {
	deltaTime := rl.GetFrameTime()
	if !g.pause {

		// implement pause logic somehow
		if rl.IsKeyPressed(rl.KeyP) {
			g.pause = !g.pause
		}

		g.player.Move()

		if rl.IsKeyPressed(rl.KeySpace) {
			projectile := player.Projectile{}
			projectile.Init(g.player.Position, rl.Vector2{X: 0, Y: -1}) // for now default upwards
			projectiles = append(projectiles, projectile)
		}

		for i := 0; i < len(projectiles); {
			projectiles[i].Move(deltaTime)
			if !projectiles[i].Spawned {
				// Remove despawned projectile
				projectiles = append(projectiles[:i], projectiles[i+1:]...)
			} else {
				i++
			}
		}
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
			int32(g.player.Size.X), rl.DarkPurple,
		)

		// Draw each projectile
		for _, proj := range projectiles {
			if proj.Spawned {
				rl.DrawCircleV(proj.Position, 5, rl.Red)
			}
		}
	}

	if g.pause {
		rl.DrawText("GAME PAUSED", constants.ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, constants.ScreenHeight/2+constants.ScreenHeight/4-40, 40, rl.Gray)
	}

	rl.EndDrawing()
}
