package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	initRaylib()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		if showingWaves {
			for i := 0; i < numPursuitWaves; i++ {
				plotFunction(p[3*i:3*i+3], wave, waveColor)
			}
		}
		if learningFromTarget {
			plotFunction(q, targetFunction, targetColor)
			learn(20)
		}
		if showingWaveSum {
			plotFunction(p, waveSum, waveSumColor)
		}
		plotFrame()
		rl.EndDrawing()
		respondToUser()

	}

	rl.CloseWindow()
}
