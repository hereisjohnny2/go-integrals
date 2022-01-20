package integrals

import (
	"testing"

	"github.com/hereisjohnny2/go-integrals/functions/polynomial"
)

func TestNewIntegral(t *testing.T) {
	f := polynomial.Polynomial{}
	integral, err := NewIntegral(0, 1, 100, &f)

	if err != nil {
		t.Errorf("Expected not to throw an error. Got %v.", err)
	}

	if integral.area != 0 {
		t.Errorf("Expected to be %v. Got %v.", 0, integral.area)
	}

	if integral.x0 != 0 {
		t.Errorf("Expected to be %v. Got %v.", 0, integral.x0)
	}

	if integral.x1 != 1 {
		t.Errorf("Expected to be %v. Got %v.", 1, integral.x1)
	}

	if integral.n != 100 {
		t.Errorf("Expected to be %v. Got %v.", 100, integral.n)
	}
}

func TestThrowNewIntegralIfInvalidLimits(t *testing.T) {
	f := polynomial.Polynomial{}
	integral, err := NewIntegral(1, 0, 100, &f)

	if err == nil {
		t.Errorf("Expected to throw an error.")
	}

	if integral.area != 0 {
		t.Errorf("Expected to be %v. Got %v.", 0, integral.area)
	}

	if integral.x0 != 0 {
		t.Errorf("Expected to be %v. Got %v.", 0, integral.x0)
	}

	if integral.x1 != 0 {
		t.Errorf("Expected to be %v. Got %v.", 0, integral.x1)
	}

	if integral.n != 0 {
		t.Errorf("Expected to be %v. Got %v.", 0, integral.n)
	}
}
