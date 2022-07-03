package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/nithinputhenveettil/brick-game-car-racing/constants"
	"github.com/nithinputhenveettil/brick-game-car-racing/game"
)

func main() {
	g := &game.Game{}
	g.Reset()

	rl.InitWindow(int32(c.Width)+int32(c.Width), int32(c.Length)+int32(c.HeadPading), "Brick Car Racing")
	rl.SetTargetFPS(120)
	for !rl.WindowShouldClose() {
		g.LitsenKeyboardEvents()
		rl.BeginDrawing()
		rl.ClearBackground(c.BgColor)
		if g.GameOver {
			g.Draw()
			rl.DrawText("Game Over!!", 10, 150, 85, rl.White)
			rl.DrawText("Press enter key to continue!", int32(c.Width)-60, int32(c.Length-60), 20, rl.White)
		} else if g.Animation.IsActive {
			g.Draw()
			g.Animation.Animate()

		} else {
			g.NextTick()
			g.CheckGameOver()
			g.Draw()
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
