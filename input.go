package main

import rl "github.com/gen2brain/raylib-go/raylib"

func respondToUser() {
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
		showingPursuitWaves = !showingPursuitWaves
	}
	if rl.IsKeyReleased(rl.KeyA) {
		showingTargetWaves = !showingTargetWaves
	}
	if rl.IsKeyReleased(rl.KeyW) {
		wholeFreqOnly = !wholeFreqOnly
	}
	if rl.IsKeyReleased(rl.KeyF) {
		showingWaveSum = !showingWaveSum
	}
	if rl.IsKeyReleased(rl.KeyUp) {
		numPursuitWaves++
		p = randomParameters(numPursuitWaves)
	}
	if rl.IsKeyReleased(rl.KeyDown) {
		if numPursuitWaves > 1 {
			numPursuitWaves--
			p = randomParameters(numPursuitWaves)
		}
	}
	if rl.IsKeyReleased(rl.KeyPeriod) {
		numTargetWaves++
		q = randomParameters(numTargetWaves)
	}
	if rl.IsKeyReleased(rl.KeyComma) {
		if numTargetWaves > 1 {
			numTargetWaves--
			q = randomParameters(numTargetWaves)
		}
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
