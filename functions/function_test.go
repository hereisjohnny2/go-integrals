package functions

import (
	"os"
	"testing"

	"github.com/hereisjohnny2/go-integrals/functions/polynomial"
)

func createPolyFunction() polynomial.Polynomial {
	return polynomial.NewPolynomial([]float64{1, 2, 2, 1})
}

func TestPlot(t *testing.T) {
	f := createPolyFunction()

	type args struct {
		f      Function
		config PlotConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should save a plot from function data",
			args: args{
				f: &f,
				config: PlotConfig{
					XLabel:   "X",
					YLabel:   "Y",
					XLimits:  []float64{-10, 10},
					YLimits:  []float64{-10, 10},
					Filename: "polyfunc.png",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Plot(tt.args.f, tt.args.config)
		})
	}

	os.Remove("polyfunc.png")
}
