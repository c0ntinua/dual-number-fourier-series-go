package main

var numWaves = int(9)
var numParameters = int(3 * numWaves)
var numData = int(21)

var rate = Real(0.0001)
var p = randomParameters()
var q = randomParameters()
var data = make([]Datum, numData)
var grad = make([]Real, 3*numWaves)
var last_loss Dual
var maxInitialWaveMagnitude = Real(0.5)
var T int = 0
var numLearnTargetSamples = 100
var targetFunction func([]Dual, Dual) Dual = waveSum
var pursuitFunction func([]Dual, Dual) Dual = waveSum
var lossPrecision = 100
