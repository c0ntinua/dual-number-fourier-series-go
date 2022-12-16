package main

type Dual struct {
	r Real
	d Real
}

func dualSum(x, y Dual) Dual {
	return Dual{x.r + y.r, x.d + y.d}
}
func dualDiff(x, y Dual) Dual {
	return Dual{x.r - y.r, x.d - y.d}
}
func dualProduct(x, y Dual) Dual {
	return Dual{x.r * y.r, y.r*x.d + x.r*y.d}
}
func dualSin(x Dual) Dual {
	return Dual{realSin(x.r), x.d * realCos(x.r)}
}
