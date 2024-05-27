package model

import (
	"reflect"
	"testing"
)

func TestNewKernel1D(t *testing.T) {
	want := Kernel1D{Values: []int{1, 1, 1}, Size: 3}
	got, err := NewKernel1D([]int{1, 1, 1})
	if err != nil {
		t.Fatalf("%v+", err)
	}
	if !reflect.DeepEqual(want.Values, got.Values) {
		t.Errorf("expected %+v; got %+v", want.Values, got.Values)
	}
	if want.Size != got.Size {
		t.Errorf("expected %d; got %d", want.Size, got.Size)
	}
}

func TestNewKernel1DError(t *testing.T) {
	_, err := NewKernel1D([]int{0, 0, 0})
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestNewKernel2D(t *testing.T) {
	want := Kernel2D{Values: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, Size: 9}
	got, err := NewKernel2D([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}})
	if err != nil {
		t.Fatalf("%v+", err)
	}
	if !reflect.DeepEqual(want.Values, got.Values) {
		t.Errorf("expected %+v; got %+v", want.Values, got.Values)
	}
	if want.Size != got.Size {
		t.Errorf("expected %d; got %d", want.Size, got.Size)
	}
}

func TestNewKernel2DError(t *testing.T) {
	_, err := NewKernel2D([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestCalcCoefficientSum1D(t *testing.T) {
	k := Kernel1D{Values: []int{1, 1, 1}, Size: 3}
	want := k.Size
	got, err := k.CalcCoefficientSum()
	if err != nil {
		t.Fatalf("%v+", err)
	}
	if want != got {
		t.Errorf("expected %d; got %d", want, got)
	}
}

func TestCalcCoefficientSum2D(t *testing.T) {
	k := Kernel2D{Values: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, Size: 9}
	want := k.Size
	got, err := k.CalcCoefficientSum()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("expected %d; got %d", want, got)
	}
}

func TestCalcMagnitude(t *testing.T) {
	tests := []struct {
		name string
		grad Gradient2D
		want uint8
	}{
		{name: "First (1,1)", grad: Gradient2D{X: 1.0, Y: 1.0}, want: 2},
		{name: "Second (10,10)", grad: Gradient2D{X: 10, Y: 10}, want: 15},
		{name: "Third (100,100)", grad: Gradient2D{X: 100, Y: 100}, want: 142},
		{name: "Fourth (1000,1000)", grad: Gradient2D{X: 1000, Y: 1000}, want: 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grad.CalcMagnitude(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcMagnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKernel2D_GetHalfKernelSize(t *testing.T) {
	tests := []struct {
		name  string
		k     Kernel2D
		want  int
		want1 int
	}{
		{
			name: "3x3 Kernel",
			k: Kernel2D{
				Size:   9,
				Values: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
				XLen:   3,
				YLen:   3,
			},
			want:  1,
			want1: 1,
		},
		{
			name: "5x5 Kernel",
			k: Kernel2D{
				Size:   25,
				Values: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
				XLen:   5,
				YLen:   5,
			},
			want:  2,
			want1: 2,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.k.GetHalfKernelSize()
			if got != tt.want {
				t.Errorf("Kernel2D.GetHalfKernelSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Kernel2D.GetHalfKernelSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
