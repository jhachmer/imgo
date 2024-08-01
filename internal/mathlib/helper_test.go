package mathlib

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAbs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0 is 0", args: args{0}, want: 0},
		{name: "1 is 1", args: args{1}, want: 1},
		{name: "-1 is 1", args: args{-1}, want: 1},
		{name: "50 is 50", args: args{50}, want: 50},
		{name: "-50 is 50", args: args{-50}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	opt := cmp.Comparer(func(x, y float64) bool {
		if x == 0 && y == 0 {
			return true
		}
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})
	type args[T NumberSigned] struct {
		values []T
		abs    bool
	}
	type testCase[T NumberSigned] struct {
		name string
		args args[T]
		want T
	}
	testsInt := []testCase[int]{
		{
			name: "[1,2,-3] | abs = 0",
			args: args[int]{
				values: []int{1, 2, -3},
				abs:    false,
			},
			want: 0,
		},
		{
			name: "[1,2,3] = 6",
			args: args[int]{
				values: []int{1, 2, -3},
				abs:    true,
			},
			want: 6,
		},
		{
			name: "[1] = 1",
			args: args[int]{
				values: []int{1},
				abs:    false,
			},
			want: 1,
		},
		{
			name: "[-1] = -1",
			args: args[int]{
				values: []int{-1},
				abs:    false,
			},
			want: -1,
		},
		{
			name: "[-1] = 1",
			args: args[int]{
				values: []int{-1},
				abs:    true,
			},
			want: 1,
		},
	}
	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.values, tt.args.abs); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
	testsFloat := []testCase[float64]{
		{
			name: "[1., 2.,-3.] & !abs = 0",
			args: args[float64]{
				values: []float64{1., 2., -3.},
				abs:    false,
			},
			want: 0.,
		},
		{
			name: "[-1.0] & abs = 0",
			args: args[float64]{
				values: []float64{-1.},
				abs:    true,
			},
			want: 1,
		},
		{
			name: "[1.0,2.,-3.] & abs = 6",
			args: args[float64]{
				values: []float64{1., 2., -3.},
				abs:    true,
			},
			want: 6,
		},
	}
	for _, tt := range testsFloat {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.values, tt.args.abs); !cmp.Equal(got, tt.want, opt) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args[T Number] struct {
		a []T
	}
	type testCase[T Number] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{name: "[1,2,3] Max = 3", args: args[int]{[]int{1, 2, 3}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {

}
