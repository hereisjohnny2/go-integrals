package functions

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type PlotConfig struct {
	XLabel   string
	YLabel   string
	XLimits  []float64
	YLimits  []float64
	Filename string
}
type Function interface {
	F(x float64) float64
	Expression() string
	GetParamters() []float64
	SetParameters([]float64) error
}

func Plot(f Function, config PlotConfig) {
	p := plot.New()

	p.Title.Text = f.Expression()
	p.X.Label.Text = config.XLabel
	p.Y.Label.Text = config.YLabel

	function := plotter.NewFunction(f.F)
	function.Color = color.RGBA{B: 255, A: 255}

	p.Add(function)

	p.X.Min = config.XLimits[0]
	p.X.Max = config.XLimits[1]
	p.Y.Min = config.YLimits[0]
	p.Y.Max = config.YLimits[1]

	if err := p.Save(4*vg.Inch, 4*vg.Inch, config.Filename); err != nil {
		panic(err)
	}
}
