package ops

import (
	"reflect"
	"testing"

	"github.com/jhachmer/imgo/internal/mathlib"
	m "github.com/jhachmer/imgo/internal/types"
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
	type args[T mathlib.Number] struct {
		s [][]T
	}
	type testCase[T mathlib.Number] struct {
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
	type args[T mathlib.Number] struct {
		s [][]T
	}
	type testCase[T mathlib.Number] struct {
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

func TestReduce(t *testing.T) {
	type args[TValue any, TResult any] struct {
		values       []TValue
		fn           func(TResult, TValue) TResult
		initialValue TResult
	}
	type testCase[TValue any, TResult any] struct {
		name string
		args args[TValue, TResult]
		want TResult
	}
	tests := []testCase[int, int]{
		{
			name: "#1",
			args: args[int, int]{
				values: []int{1, 2, 3},
				fn: func(result int, value int) int {
					return result + value
				},
				initialValue: 0,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.values, tt.args.fn, tt.args.initialValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[TValue any, TResult any] struct {
		values []TValue
		fn     func(TValue) TResult
	}
	type testCase[TValue any, TResult any] struct {
		name string
		args args[TValue, TResult]
		want []TResult
	}
	tests := []testCase[int, int]{
		{
			name: "#1",
			args: args[int, int]{
				values: []int{1, 2, 3},
				fn: func(i int) int {
					return i * i
				},
			},
			want: []int{1, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.values, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransposeMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3x3",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "2x3",
			args: args{
				matrix: [][]int{{1, 2}, {4, 5}, {7, 8}},
			},
			want: [][]int{{1, 4, 7}, {2, 5, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransposeMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransposeMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		values []int
		fn     func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "#1",
			args: args{values: []int{-1, -2, -3, 1, 2, 3},
				fn: func(i int) bool {
					if i >= 0 {
						return true
					}
					return false
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.values, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
