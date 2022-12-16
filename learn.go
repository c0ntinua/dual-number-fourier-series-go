package main

func learn(n int) {
	for i := 0; i < n; i++ {
		gradient := gradient(p)
		for j := range p {
			p[j].r -= rate * gradient[j]
		}
	}
}
