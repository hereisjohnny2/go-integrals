package exponential

import (
	"errors"
	"fmt"
	"math"
)

// Express an Exponential Function
// f(x) = a*e^(b*x)
type Exponential struct {
	a float64
	b float64
}

func NewExponential(_params []float64) (Exponential, error) {
	if len(_params) >= 2 {
		return Exponential{a: _params[0], b: _params[1]}, nil
	}
	return Exponential{}, errors.New("params slice is too short (should have at least 2 elements)")
}

func (f Exponential) F(x float64) float64 {
	return f.a * math.Exp(f.b*x)
}

func (f Exponential) Expression() string {
	return fmt.Sprintf("f(x)=%v*e^(%v*x)", f.a, f.b)
}

func (f Exponential) GetParamters() []float64 {
	return []float64{f.a, f.b}
}

func (f *Exponential) SetParameters(_params []float64) error {
	if len(_params) >= 2 {
		f.a = _params[0]
		f.b = _params[1]
		return nil
	}
	return errors.New("params slice is too short (should have at least 2 elements)")
}
