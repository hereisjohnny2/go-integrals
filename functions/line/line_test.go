package line

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/hereisjohnny2/go-integrals/functions"
)

func makeListPoint() []Point {
	return []Point{
		{1, 2},
		{2, 4},
	}
}

func makeLine() Line {
	return Line{
		points: makeListPoint(),
	}
}

func createTestFile() {
	os.Create("test.txt")

	ds := ""
	for _, p := range makeListPoint() {
		x := fmt.Sprintf("%v", p.x)
		y := fmt.Sprintf("%v", p.y)

		s := strings.Join([]string{x, y}, "\t")
		ds += s + "\n"
	}

	if err := ioutil.WriteFile("test.txt", []byte(ds), 0644); err != nil {
		panic(err)
	}
}

func TestNewLine(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			name: "should be able to create a new line from a list o points",
			args: args{points: makeListPoint()},
			want: makeLine(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLine(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_Length(t *testing.T) {
	type fields struct {
		points []Point
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "should be able give the full length of a line",
			fields: fields{points: makeListPoint()},
			want:   math.Sqrt(math.Pow(4-2, 2) + math.Pow(2-1, 2)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				points: tt.fields.points,
			}
			if got := l.Length(); got != tt.want {
				t.Errorf("Line.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	createTestFile()

	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			name: "should be able to load a list of point from a file",
			args: args{filename: "test.txt"},
			want: makeLine(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Load(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove("test.txt")
}

func TestLine_Save(t *testing.T) {
	type fields struct {
		points []Point
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "should be able to save a line data in a file",
			fields: fields{makeListPoint()},
			args:   args{"test.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				points: tt.fields.points,
			}
			l.Save(tt.args.filename)
		})
	}

	os.Remove("test.txt")
}

func TestLine_Plot(t *testing.T) {
	type fields struct {
		points []Point
	}
	type args struct {
		config functions.PlotConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "should save a plot image from the line data",
			fields: fields{makeListPoint()},
			args: args{functions.PlotConfig{
				Filename: "test_plot.png",
				XLabel:   "X",
				YLabel:   "Y",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				points: tt.fields.points,
			}
			l.Plot(tt.args.config)
		})
	}

	os.Remove("test_plot.png")
}

func TestLine_GetPoints(t *testing.T) {
	type fields struct {
		points []Point
	}
	tests := []struct {
		name   string
		fields fields
		want   []Point
	}{
		{
			name:   "should return the list of points",
			fields: fields{makeListPoint()},
			want:   makeListPoint(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				points: tt.fields.points,
			}
			if got := l.GetPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Line.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
