package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	initRaylib()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		if showingPursuitWaves {
			for i := 0; i < numPursuitWaves; i++ {
				plotFunction(p[3*i:3*i+3], wave, pursuitWaveColor)
			}
		}
		if showingTargetWaves {
			for i := 0; i < numTargetWaves; i++ {
				plotFunction(q[3*i:3*i+3], wave, targetWaveColor)
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
