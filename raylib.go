package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func initRaylib() {
	setWaveColors()
	rl.InitWindow(screenWidth, screenHeight, "dual-sign-go")
	rl.SetTargetFPS(60)
}
