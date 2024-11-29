package main

import (
	"gengine/constants"
	"gengine/enemy"
	"gengine/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	pause  bool
	player player.Player
	enemy  enemy.Enemy
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.player.Spawn()
	g.enemy.Spawn()
}

var projectiles []player.Projectile

func (g *Game) Update() {
	deltaTime := rl.GetFrameTime()
	if !g.pause {

		// implement pause logic somehow
		if rl.IsKeyPressed(rl.KeyP) {
			g.pause = !g.pause
		}

		g.player.Update()

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
		g.player.Draw()
		g.enemy.Draw()

		// Draw each projectile
		for _, proj := range projectiles {
			if proj.Spawned {
				proj.Draw()
			}
		}
	}

	if g.pause {
		rl.DrawText("GAME PAUSED", constants.ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, constants.ScreenHeight/2+constants.ScreenHeight/4-40, 40, rl.Gray)
	}

	rl.EndDrawing()
}
