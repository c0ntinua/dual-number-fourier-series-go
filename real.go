package main

import (
	"math"
	"math/rand"
)

type Real float64

func randomReal(a, b Real) Real {
	return a + (b-a)*Real(rand.Float64())
}
func realSin(x Real) Real {
	return Real(math.Sin(float64(x)))
}
func realCos(x Real) Real {
	return Real(math.Cos(float64(x)))
}
func realFrac(x Real) Real {
	_, frac := math.Modf(float64(x))
	return Real(frac)
}
