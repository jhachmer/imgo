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
