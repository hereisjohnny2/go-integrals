package polynomial

import (
	"fmt"
	"math"
)

// f(x) = p0 + p1*x + p2*x^2 + ... + p(n-1)*x^(n-1)
type Polynomial struct {
	params []float64
}

func NewPolynomial(_params []float64) Polynomial {
	return Polynomial{params: _params}
}

func (f Polynomial) F(x float64) float64 {
	y := 0.0
	for i, elm := range f.params {
		y += elm * math.Pow(x, float64(i))
	}

	return y
}

func (f Polynomial) Expression() string {
	expr := "f(x)="
	for i, elm := range f.params {
		elmStr := fmt.Sprintf("%v", elm)
		if i == 0 {
			expr += elmStr
		} else {
			indexStr := fmt.Sprintf("%v", i)
			if elm > 0 {
				expr += "+"
			}
			expr += elmStr + "x^" + indexStr
			i++
		}
	}
	return expr
}

func (f Polynomial) GetParamters() []float64 {
	return f.params
}

func (f *Polynomial) SetParameters(params []float64) {
	f.params = params
}
