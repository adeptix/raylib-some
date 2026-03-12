package main

// add this to go.mod before compilation
// replace github.com/gen2brain/raylib-go/raylib => ./Raylib-Go-Wasm/raylib

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable)
	rl.InitWindow(0, 0, "raylib [core] example - basic window")

	var update = func() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.SetMainLoop(update)

	for !rl.WindowShouldClose() {
		update()
	}
	rl.CloseWindow()
}
