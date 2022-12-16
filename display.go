package main

import (
	"image/color"
	"math"
)

var plotXMin = Real(-math.Pi)
var plotXMax = Real(math.Pi)
var plotYMin = Real(-math.Pi)
var plotYMax = Real(math.Pi)
var dataXMin = Real(-math.Pi)
var dataXMax = Real(math.Pi)
var dataYMin = Real(-2.5)
var dataYMax = Real(2.5)

var dataRadius = Real(10)
var modelThickness = Real(10)
var waveThickness = Real(10)
var targetThickness = Real(10)
var frameThickness = float32(10)
var waveHue = uint8(40)

var waveColors = make([]color.RGBA, numPursuitWaves)
var targetColor = color.RGBA{0, 0, 255, 255}
var waveSumColor = color.RGBA{255, 0, 0, 255}
var waveColor = color.RGBA{waveHue, waveHue, waveHue, 255}
var frameColor = color.RGBA{255, 255, 255, 255}
var backgroundColor = color.RGBA{0, 0, 0, 255}

var hppu = Real(screenWidth / (plotXMax - plotXMin))
var vppu = Real(screenHeight / (plotYMax - plotYMin))

var numPlotSegments = 500
var plotSegmentGap = (plotXMax - plotXMin) / Real(numPlotSegments)
