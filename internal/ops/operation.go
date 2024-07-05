package ops

import (
	"slices"

	"github.com/jhachmer/imgo/internal/mathlib"
	"github.com/jhachmer/imgo/internal/types"
)

func ClampPixel[T mathlib.Number](value T) uint8 {
	var result uint8
	if value < 0 {
		result = 0
	} else if value > 255 {
		result = 255
	} else {
		result = uint8(value)
	}
	return result
}

func GenerateSlice[T any](x, y int) [][]T {
	res := make([][]T, y)
	for i := 0; i < y; i++ {
		res[i] = make([]T, x)
	}
	return res
}

func GenerateComplexSlice(pixels [][]uint8) [][]types.Complex {
	c := make([][]types.Complex, len(pixels))
	for i := range c {
		c[i] = make([]types.Complex, len(pixels[i]))
	}
	for j := range len(pixels) {
		for i := range len(pixels[j]) {
			c[j][i] = *types.NewComplex(float64(pixels[j][i]), 0)
		}
	}

	return c
}

func TransposeMatrix[T any](matrix [][]T) [][]T {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]T, cols)
	for i := range transposed {
		transposed[i] = make([]T, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

func FindMaxIn2DSlice[T mathlib.Number](s [][]T) T {
	var subMax, curMax T
	for i := range s {
		subMax = slices.Max(s[i])
		if subMax > curMax {
			curMax = subMax
		}
	}
	return curMax
}

func Reduce[TValue, TResult any](values []TValue, fn func(TResult, TValue) TResult, initialValue TResult) TResult {
	result := initialValue
	for _, value := range values {
		result = fn(result, value)
	}
	return result
}

func Map[TValue, TResult any](values []TValue, fn func(TValue) TResult) []TResult {
	result := make([]TResult, len(values))

	for i, value := range values {
		result[i] = fn(value)
	}
	return result
}

func Filter[TValue any](values []TValue, fn func(TValue) bool) []TValue {
	result := make([]TValue, 0)

	for _, value := range values {
		if ok := fn(value); ok {
			result = append(result, value)
		}
	}
	return result
}
