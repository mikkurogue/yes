package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 640, "raylib [core] exacmple window")

	game := NewGame()
	game.pause = false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		game.Update()

		game.Draw()
	}

	rl.CloseWindow()
}
