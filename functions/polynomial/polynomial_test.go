package polynomial

import (
	"math"
	"testing"
)

func makeMockParams() []float64 {
	return []float64{1, 2, 3, 4}
}

func TestNewPolynomial(t *testing.T) {
	mockedParams := makeMockParams()
	params := NewPolynomial(mockedParams).params

	if params == nil {
		t.Errorf("Expected params to be not null.")
	}

	for key, elm := range params {
		if elm != mockedParams[key] {
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
	poly := NewPolynomial(makeMockParams())

	expr := poly.Expression()
	expected := "f(x)=1+2x^1+3x^2+4x^3"

	if expr != expected {
		t.Errorf("Expected to be %v. Got %v.", expected, expr)
	}
}

func TestGetParams(t *testing.T) {
	params := NewPolynomial(makeMockParams()).GetParamters()
	expected := makeMockParams()

	if params[0] != expected[0] {
		t.Errorf("Expected to be %v. Got %v.", params, expected)
	}
}

func TestSetParams(t *testing.T) {
	poly := NewPolynomial(makeMockParams())
	expected := []float64{5, 5}

	poly.SetParameters(expected)
	params := poly.params

	if params[0] != expected[0] && params[1] != expected[1] {
		t.Errorf("Expected to be %v. Got %v.", params, expected)
	}
}
