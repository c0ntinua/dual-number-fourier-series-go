package main

func loss(p []Dual) Dual {
	var loss Dual
	for i := 0; i < numLearnTargetSamples; i++ {
		x := Dual{randomReal(plotXMin, plotXMax), 0}
		diff := dualDiff(pursuitFunction(p, x), targetFunction(q, x))
		sqr_diff := dualProduct(diff, diff)
		loss = dualSum(loss, sqr_diff)
	}
	return loss
}
