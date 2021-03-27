package main

type xy struct {
	x []float64
	y []float64
}

var (
	a = 0.9463
	b = -0.3985
)

func main() {
	a, b = reg()
	plotLR(a, b)
}
