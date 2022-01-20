package main

import (
	"fmt"

	"github.com/hereisjohnny2/go-integrals/functions"
	"github.com/hereisjohnny2/go-integrals/functions/polynomial"
)

func main() {
	poly := polynomial.NewPolynomial([]float64{1, 2, 2, 1})

	fmt.Println(poly.Expression())

	functions.Plot(&poly, functions.PlotConfig{
		XLabel:   "X",
		YLabel:   "Y",
		XLimits:  []float64{-10, 10},
		YLimits:  []float64{-10, 10},
		Filename: "polyfunc.png",
	})
}
