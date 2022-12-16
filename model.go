package main

func waveSum(p []Dual, x Dual) Dual {
	var sum Dual
	for i := 0; i < numWaves; i++ {
		sum = dualSum(sum, wave(p[3*i:3*i+3], x))
	}
	return sum
}

func wave(p []Dual, x Dual) Dual {
	return dualProduct(p[0], dualSin(dualSum(p[1], dualProduct(p[2], x))))
}
