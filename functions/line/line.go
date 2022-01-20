package line

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/hereisjohnny2/go-integrals/functions"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Line struct {
	points []Point
}

func NewLine(points []Point) Line {
	return Line{points}
}

func (l Line) Length() float64 {
	length := 0.0
	for i := 1; i < len(l.points); i++ {
		length += math.Sqrt(math.Pow((l.points[i].y-l.points[i-1].y), 2) + math.Pow((l.points[i].x-l.points[i-1].x), 2))
	}

	return length
}

func Load(filename string) Line {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	ds := strings.Split(string(bs), "\n")

	points := []Point{}
	for _, p := range ds {
		s := strings.Split(p, "\t")
		if len(s) == 2 {
			x, _ := strconv.ParseFloat(s[0], 64)
			y, _ := strconv.ParseFloat(s[1], 64)
			points = append(points, Point{x, y})
		}
	}

	return Line{points}
}

func (l Line) Save(filename string) {
	ds := ""
	for _, p := range l.points {
		x := fmt.Sprintf("%v", p.x)
		y := fmt.Sprintf("%v", p.y)

		s := strings.Join([]string{x, y}, "\t")
		ds += s + "\n"
	}

	if err := ioutil.WriteFile(filename, []byte(ds), 0644); err != nil {
		panic(err)
	}
}

func (l Line) Plot(config functions.PlotConfig) {
	p := plot.New()

	pts := make(plotter.XYs, len(l.points))

	for i, point := range l.points {
		pts[i].X = point.x
		pts[i].Y = point.y
	}

	if err := plotutil.AddLinePoints(p, pts); err != nil {
		panic(err)
	}

	p.Title.Text = "Plot from Data"
	p.X.Label.Text = config.XLabel
	p.Y.Label.Text = config.YLabel

	if err := p.Save(4*vg.Inch, 4*vg.Inch, config.Filename); err != nil {
		panic(err)
	}
}

func (l Line) GetPoints() []Point {
	return l.points
}

func (l Line) CalculateArea() float64 {
	area := 0.0
	for i := 1; i < len(l.points); i++ {
		a := l.points[i-1].y
		b := l.points[i].y
		h := l.points[i].x - l.points[i-1].x
		area += (a + b) * h / 2
	}
	return area
}
