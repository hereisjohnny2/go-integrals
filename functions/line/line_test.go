package line

import (
	"reflect"
	"testing"
)

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
			name: "should be able to create a new line with a slice of points",
			args: args{
				points: []Point{{1, 2}, {3, 4}},
			},
			want: Line{points: []Point{{1, 2}, {3, 4}}},
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
	tests := []struct {
		name string
		l    Line
		want float64
	}{
		{
			name: "should be able to create a new line with a slice of points",
			args: Line{points: []Point{{1, 2}, {3, 4}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Length(); got != tt.want {
				t.Errorf("Line.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
