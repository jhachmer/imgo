package ops

import (
	m "github.com/jhachmer/imgo/types"
	"reflect"
	"testing"
)

func ComplexSlice() [][]m.Complex {
	return [][]m.Complex{
		{
			{Re: 1, Im: 0}, {Re: 2, Im: 0},
		},
		{
			{Re: 3, Im: 0}, {Re: 4, Im: 0},
		},
	}
}

func TestGenerateComplexSlice(t *testing.T) {
	type args struct {
		pixelLine [][]uint8
	}
	tests := []struct {
		name string
		args args
		want [][]m.Complex
	}{
		{
			name: "2",
			args: args{[][]uint8{{1, 2}, {3, 4}}},
			want: ComplexSlice(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateComplexSlice(tt.args.pixelLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateComplexSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxIn2DSliceInt(t *testing.T) {
	type args[T Number] struct {
		s [][]T
	}
	type testCase[T Number] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{name: "3x3", args: args[int]{s: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, want: 9},
		{name: "3x3", args: args[int]{s: [][]int{{0, 0, 0}, {0, 5, 0}, {0, 0, 0}}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxIn2DSlice(tt.args.s); got != tt.want {
				t.Errorf("FindMaxIn2DSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxIn2DSliceFloat(t *testing.T) {
	type args[T Number] struct {
		s [][]T
	}
	type testCase[T Number] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[float64]{
		{name: "3x3", args: args[float64]{s: [][]float64{{1.1, 2.2, 3.3}, {4.4, 5.5, 6.6}, {7.7, 8.8, 9.9}}}, want: 9.9},
		{name: "3x3", args: args[float64]{s: [][]float64{{0.1, 0.1, 0.1}, {0.2, 2.222, 2.2223}, {0.6, 2.1, 1.55754334}}}, want: 2.2223},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxIn2DSlice(tt.args.s); got != tt.want {
				t.Errorf("FindMaxIn2DSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
