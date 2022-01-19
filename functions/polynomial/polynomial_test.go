package polynomial

import (
	"math"
	"testing"
)

func makeMockParams() map[string]float64 {
	return map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
}

func TestNewPolynomial(t *testing.T) {
	mockedParams := makeMockParams()
	params := NewPolynomial(mockedParams).params

	if params == nil {
		t.Errorf("Expected params to be not null.")
	}

	for key := range params {
		if params[key] != mockedParams[key] {
			t.Errorf("Expected to be %v. Got %v.", mockedParams, params)
			break
		}
	}
}

func TestF(t *testing.T) {
	f := NewPolynomial(makeMockParams())

	y := f.F(5)
	expected := 1 + 2*5 + 3*math.Pow(5, 2) + 4*math.Pow(5, 3)

	if y != expected {
		t.Errorf("Expected to be %v. Got %v.", expected, y)
	}
}

func TestExpression(t *testing.T) {
	expr := NewPolynomial(makeMockParams()).Expression()
	expected := "f(x)=1+2x^1+3x^2+4x^3"

	if expr != expected {
		t.Errorf("Expected to be %v. Got %v.", expected, expr)
	}
}
