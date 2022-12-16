package main

import (
	"image/color"
	"math/rand"
)

func setWaveColors() {
	for i := 0; i < numPursuitWaves; i++ {
		waveColors[i] = color.RGBA{uint8(rand.Uint32() % 256), uint8(rand.Uint32() % 256), uint8(rand.Uint32() % 256), 255}
		waveColors[i] = color.RGBA{155, 155, 155, 255}
		//waveColors[i] = color.RGBA{255, 0, 0, 255}
	}

}
