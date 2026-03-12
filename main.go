package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"raylib-some/snowflake"
)

func main() {
	closeFunc := snowflake.InitWindow()
	defer closeFunc()

	for !rl.WindowShouldClose() {
		snowflake.UpdateFunc()
	}
}
