package kernel

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

func TestNewKernel2DEdge(t *testing.T) {
	want := Kernel2D{Values: [][]int{{0, 1, 0}, {1, -4, 1}, {0, 1, 0}}, Size: 0}
	got, err := NewKernel2D([][]int{{0, 1, 0}, {1, -4, 1}, {0, 1, 0}})
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
	_, err := NewKernel2D([][]int{{0, 0}, {0, 0, 0}, {0, 0}})
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestCalcCoefficientSum1D(t *testing.T) {
	k := Kernel1D{Values: []int{1, 1, 1}, Size: 3}
	want := k.Size
	err := k.CalcCoefficientSum()
	if err != nil {
		t.Fatalf("%v+", err)
	}
	if want != k.Size {
		t.Errorf("expected %d; got %d", want, k.Size)
	}
}

func TestCalcCoefficientSum2D(t *testing.T) {
	k := Kernel2D{Values: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, Size: 9}
	want := k.Size
	k.CalcCoefficientSum()
	if want != k.Size {
		t.Errorf("expected %d; got %d", want, k.Size)
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
				Width:  3,
				Height: 3,
			},
			want:  1,
			want1: 1,
		},
		{
			name: "5x5 Kernel",
			k: Kernel2D{
				Size:   25,
				Values: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
				Width:  5,
				Height: 5,
			},
			want:  2,
			want1: 2,
		},
		{
			name: "7x7 Kernel",
			k: Kernel2D{
				Size:   25,
				Values: [][]int{{1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}},
				Width:  7,
				Height: 7,
			},
			want:  3,
			want1: 3,
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
