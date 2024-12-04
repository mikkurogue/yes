package enemy

import (
	"fmt"
	"gengine/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	Spawned        bool
	Health         float32
	Position, Size rl.Vector2
}

func (e *Enemy) Update() {
	fmt.Println("NYI: ENEMY UPDATE FUNC")
}

func (e *Enemy) Draw() {
	rl.DrawCircleV(
		e.Position,
		25,
		rl.DarkGreen,
	)
}

func (e *Enemy) Spawn() {
	e.Spawned = true
	e.Health = 100.0

	e.Position = rl.Vector2{
		X: float32(constants.ScreenWidth / 2),
		Y: float32(constants.ScreenHeight * 7 / 16),
	}
	e.Size = rl.Vector2{
		X: float32(constants.ScreenWidth / 10),
		Y: 20.0,
	}
}

func (e *Enemy) TakeDamage(amount float32) {
	e.Health -= amount

	if e.Health <= 0 {
		e.Health = 0
		e.Spawned = false
	}
}
