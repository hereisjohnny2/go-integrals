package integrals

import "github.com/hereisjohnny2/go-integrals/functions"

type Integral interface {
	Area(x0 float64, x1 float64, n int, f functions.Function) float64
}
