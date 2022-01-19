package exponential

import (
	"math"
	"testing"
)

func makeMockParams() []float64 {
	return []float64{1, 2}
}

func TestNewExponential(t *testing.T) {
	expected := makeMockParams()
	f, err := NewExponential(expected)

	if err != nil {
		t.Errorf("Expect to not throw error. Got %v", err)
	}

	if f.a != expected[0] {
		t.Errorf("Expected to be %v. Got %v.", f.a, expected)
	}

	if f.b != expected[1] {
		t.Errorf("Expected to be %v. Got %v.", f.b, expected)
	}
}

func TestNotNewExponentialIfSliceToShort(t *testing.T) {
	_, err := NewExponential([]float64{1})

	if err == nil {
		t.Errorf("Expected to throw error.")
	}
}

func TestF(t *testing.T) {
	f, _ := NewExponential(makeMockParams())

	y := f.F(5)
	expected := 1 * math.Exp(2*5)
	if y != expected {
		t.Errorf("Expected to be %v. Got %v.", expected, y)
	}
}

func TestExpression(t *testing.T) {
	f, _ := NewExponential(makeMockParams())

	expr := f.Expression()
	expected := "f(x)=1*e^(2*x)"

	if expr != expected {
		t.Errorf("Expected to be %v. Got %v.", expected, expr)
	}
}

func TestGetParams(t *testing.T) {
	f, _ := NewExponential(makeMockParams())

	expected := makeMockParams()

	if f.a != expected[0] {
		t.Errorf("Expected to be %v. Got %v.", f.a, expected)
	}

	if f.b != expected[1] {
		t.Errorf("Expected to be %v. Got %v.", f.b, expected)
	}
}

func TestSetParams(t *testing.T) {
	f, _ := NewExponential(makeMockParams())
	expected := []float64{5, 5}

	err := f.SetParameters(expected)

	if err != nil {
		t.Errorf("Expect to not throw error. Got %v", err)
	}

	if f.a != expected[0] {
		t.Errorf("Expected to be %v. Got %v.", f.a, expected)
	}

	if f.b != expected[1] {
		t.Errorf("Expected to be %v. Got %v.", f.b, expected)
	}
}

func TestNotSetParamsIfSliceIsTooShort(t *testing.T) {
	f, _ := NewExponential(makeMockParams())

	err := f.SetParameters([]float64{1})

	if err == nil {
		t.Errorf("Expected to throw error.")
	}
}
