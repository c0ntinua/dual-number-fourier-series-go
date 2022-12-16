package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func scaledRealX(x Real) float32 {
	return float32(hppu * (x - plotXMin))
}

func scaledRealY(y Real) float32 {
	return float32(screenHeight - vppu*(y-plotYMin))
}

func plotFunction(p []Dual, f func([]Dual, Dual) Dual, color rl.Color) {
	firstX := scaledRealX(plotXMin)
	firstY := scaledRealY(f(p, Dual{plotXMin, 0}).r)
	var lastPoint = rl.Vector2{X: firstX, Y: firstY}
	var thisPoint rl.Vector2
	for j := 1; j < numPlotSegments; j++ {
		arg := dataXMin + Real(j)*plotSegmentGap
		x := scaledRealX(arg)
		y := scaledRealY(f(p, Dual{arg, 0}).r)
		thisPoint = rl.Vector2{X: x, Y: y}
		rl.DrawLineEx(lastPoint, thisPoint, float32(modelThickness), color)
		lastPoint = thisPoint
	}
}
func plotFrame() {
	rl.DrawRectangleLinesEx(rl.Rectangle{0, 0, screenWidth, screenHeight}, frameThickness, frameColor)
}
