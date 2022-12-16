package main

var numPursuitWaves = int(6)
var numTargetWaves = int(9)

var numData = int(21)

var rate = Real(0.0001)
var p = randomParameters(numPursuitWaves)
var q = randomParameters(numTargetWaves)
var grad = make([]Real, 3*numPursuitWaves)
var last_loss Dual
var maxInitialWaveMagnitude = Real(0.5)
var T int = 0
var numLearnTargetSamples = 100
var targetFunction func([]Dual, Dual) Dual = waveSum
var pursuitFunction func([]Dual, Dual) Dual = waveSum
var lossPrecision = 100
