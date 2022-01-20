package line

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Line struct {
	points []Point
}

func NewLine(points []Point) Line {
	return Line{points}
}

func (l Line) Length() float64 {
	len := 0.0
	for i := range l.points[1:] {
		len += math.Sqrt(math.Pow((l.points[i].y-l.points[i-1].y), 2) + math.Pow((l.points[i].x-l.points[i-1].x), 2))
	}

	return len
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
		x, _ := strconv.ParseFloat(s[0], 64)
		y, _ := strconv.ParseFloat(s[1], 64)
		points = append(points, Point{x, y})
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

	ioutil.WriteFile(filename, []byte(ds), 0644)
}
