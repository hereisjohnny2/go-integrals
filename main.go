package main

import (
	"fmt"

	"github.com/hereisjohnny2/go-integrals/functions"
	"github.com/hereisjohnny2/go-integrals/functions/polynomial"
)

func useInterface(f functions.Function) {
	f.SetParameters([]float64{1, 2, 3})
}

func main() {
	poly := polynomial.NewPolynomial([]float64{7, 8, 9})
	useInterface(&poly)

	fmt.Println(poly.Expression())
}
