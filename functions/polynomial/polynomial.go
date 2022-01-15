package polynomial

import (
	"math"
	"strconv"
)

// f(x) = p0 + p1*x + p2*x^2 + ... + p(n-1)*x^(n-1)
type Polynomial struct {
	params map[string]float64
}

func NewPolynomial(_params map[string]float64) Polynomial {
	return Polynomial{params: _params}
}

func (f Polynomial) F(x float64) float64 {
	y := 0.0
	i := 0
	for _, elm := range f.params {
		if i == 0 {
			y += elm
		} else {
			y += elm * math.Pow(x, float64(i))
		}
		i++
	}

	return y
}

func (f Polynomial) Expression() string {
	expr := "f(x)="
	for key, elm := range f.params {
		if key == "a" {
			expr += strconv.FormatFloat(elm, 'E', -1, 64)
		} else {
			if elm > 0 {
				expr += "+" + strconv.FormatFloat(elm, 'E', -1, 64)
			} else if elm < 0 {
				expr += strconv.FormatFloat(elm, 'E', -1, 64)
			}
		}
	}
	return expr
}

func (f Polynomial) GetParamters() map[string]float64 {
	return f.params
}

func (f *Polynomial) SetParameters(params map[string]float64) {
	f.params = params
}
