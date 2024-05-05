package model

import (
	"reflect"
	"testing"
)

func TestNewKernel1D(t *testing.T) {
	want := Kernel1D{Values: []int{1, 1, 1}, Size: 3}
	got := NewKernel1D([]int{1, 1, 1})
	if !reflect.DeepEqual(want.Values, *&got.Values) {
		t.Errorf("expected %+v; got %+v", want.Values, got.Values)
	}
	if want.Size != got.Size {
		t.Errorf("expected %d; got %d", want.Size, got.Size)
	}
}

func TestNewKernel2D(t *testing.T) {
  want := Kernel2D{Values: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, Size: 9}
	got := NewKernel2D([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}})
	if !reflect.DeepEqual(want.Values, got.Values) {
		t.Errorf("expected %+v; got %+v", want.Values, got.Values)
	}
	if want.Size != got.Size {
		t.Errorf("expected %d; got %d", want.Size, got.Size)
	}
}

func TestCalcCoeffSum1D(t *testing.T) {
	k := Kernel1D{Values: []int{1, 1, 1}, Size: 3}
	want := k.Size
	got := k.CalcCoeffSum()
	if want != got {
		t.Errorf("expected %d; got %d", want, got)
	}
}

func TestCalcCoeffSum2D(t *testing.T) {
  k := Kernel2D{Values: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, Size: 9}
  want := k.Size
  got := k.CalcCoeffSum()
  if want != got {
		t.Errorf("expected %d; got %d", want, got)
  }
}

func TestCalcMagnitude(t *testing.T) {
	type test struct {
		name string
		args Gradient2D
		want uint8
	}
  tests := []test{
    {name: "First (1,1)", args: Gradient2D{X: 1 + 127, Y: 1 + 127}, want: 0 },
    {name: "Second (10,10)", args: Gradient2D{X: 10 + 127, Y: 10 + 127}, want: 1},
    {name: "Third (100,100)", args: Gradient2D{X: 100 + 127, Y: 100 + 127}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.CalcMagnitude(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcMagnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
