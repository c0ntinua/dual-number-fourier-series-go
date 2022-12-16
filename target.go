package main

import "math"

func targetSin(p []Dual, x Dual) Dual {
	return dualProduct(Dual{2, 0}, dualSin(x))
}

func sawWave(p []Dual, x Dual) Dual {
	i, frac := math.Modf(float64(x.r))
	return Dual{Real(2 * frac * math.Pow(-1, float64(int(i)%2))), 0}
}
