package main

import (
	"math"
)

func randomParameters(numberOfWaves int) []Dual {
	p := make([]Dual, 3*numberOfWaves)
	for i := 0; i < numberOfWaves; i++ {
		p[3*i+0] = Dual{randomReal(-maxInitialWaveMagnitude, -maxInitialWaveMagnitude), 0}
		p[3*i+1] = Dual{randomReal(-math.Pi, math.Pi), 0}
		p[3*i+2] = Dual{Real(i + 1), 0}
	}
	return p
}
