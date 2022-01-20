package integrals

import (
	"errors"

	"github.com/hereisjohnny2/go-integrals/functions"
)

type Integral struct {
	area float64
	x0   float64
	x1   float64
	n    int
	f    functions.Function
}

func NewIntegral(x0 float64, x1 float64, n int, f functions.Function) (Integral, error) {
	if x1 < x0 {
		return Integral{}, errors.New("expected x1 to be bigger than x0")
	}

	return Integral{0, x0, x1, n, f}, nil
}

// Implements the Trapezoid rule:
//
// Area = h[(f(a) + f(b)) / 2 + sum_{1}_{N-1}(f(x_i))]
func (it *Integral) Trapezoidal(x0 float64, x1 float64, n int, f functions.Function) float64 {
	h := (x1 - x0) / float64(n)
	it.area = (f.F(x0) + f.F(x1)) / 2

	for i := 1; i < n; i++ {
		it.area += f.F(x0 + h*float64(i))
	}

	return it.area * h
}
