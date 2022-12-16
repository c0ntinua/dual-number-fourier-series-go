package main

func gradient(p []Dual) []Real {
	grad := make([]Real, len(p))
	for i := range p {
		p[i].d = 1.0
		grad[i] = loss(p).d
		p[i].d = 0.0
	}
	return grad
}
