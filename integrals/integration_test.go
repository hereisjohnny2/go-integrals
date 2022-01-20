package integrals

import (
	"testing"

	"github.com/hereisjohnny2/go-integrals/functions"
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

func TestIntegral_Trapezoidal(t *testing.T) {
	poly := polynomial.NewPolynomial([]float64{2, 2, 2})

	type fields struct {
		area float64
		x0   float64
		x1   float64
		n    int
		f    functions.Function
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Polynomial function",
			fields: fields{
				area: 0,
				x0:   0,
				x1:   1,
				n:    100,
				f:    &poly,
			},
			want: 3.6667,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &Integral{
				area: tt.fields.area,
				x0:   tt.fields.x0,
				x1:   tt.fields.x1,
				n:    tt.fields.n,
				f:    tt.fields.f,
			}
			if got := it.Trapezoidal(); (got - tt.want) > 1e-6 {
				t.Errorf("Integral.Trapezoidal() = %v, want %v", got, tt.want)
				t.Errorf("Integral.Trapezoidal() error = %v, want %v", got-tt.want, 1e-6)
			}
		})
	}
}
