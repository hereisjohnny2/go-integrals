package trapezoid

import "github.com/hereisjohnny2/go-integrals/functions"

type Trapezoid struct {
	area float64
}

func (it *Trapezoid) Area(x0 float64, x1 float64, n int, f functions.Function) float64 {
	h := (x1 - x0) / float64(n)
	it.area = h / 2 * (f.F(x0) + f.F(x1))

	for i := 1; i < n; i++ {
		it.area += f.F(x0+h*float64(i)) * h
	}

	return it.area
}
