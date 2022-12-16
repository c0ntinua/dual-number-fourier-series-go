package main

import rl "github.com/gen2brain/raylib-go/raylib"

func respondToUser() {
	if rl.IsKeyReleased(rl.KeyD) {
		randomizeData()
	}
	if rl.IsKeyReleased(rl.KeyP) {
		p = randomParameters(numPursuitWaves)
	}
	if rl.IsKeyReleased(rl.KeyT) {
		q = randomParameters(numTargetWaves)
	}
	if rl.IsKeyReleased(rl.KeyQ) {
		rl.CloseWindow()
	}
	if rl.IsKeyReleased(rl.KeyS) {
		showingWaves = !showingWaves
	}
	if rl.IsKeyReleased(rl.KeyF) {
		showingWaveSum = !showingWaveSum
	}
	if rl.IsKeyReleased(rl.KeyRight) {
		rate += 0.0001
	}
	if rl.IsKeyReleased(rl.KeyLeft) {
		rate -= 0.0001
	}
	if rl.IsKeyReleased(rl.KeyZero) {
		rate = 0.000
	}
	if rl.IsKeyReleased(rl.KeyL) {
		learningFromTarget = !learningFromTarget
	}
}
