package util

import (
	"github.com/jhachmer/imgo/model"
)

func ClampPixel(value, upper, lower int) int {
	if value < lower {
		value = lower
	}
	if value > upper {
		value = upper
	}
	return value
}

func GeneratePixelSlice(x, y int) [][]uint8 {
	res := make([][]uint8, y)
	for i := 0; i < y; i++ {
		res[i] = make([]uint8, x)
	}
	return res
}

func GenerateComplexSlice(pixels [][]uint8) [][]model.Complex {
	c := make([][]model.Complex, len(pixels))
	for i := range c {
		c[i] = make([]model.Complex, len(pixels[i]))
	}
	for j := range len(pixels) {
		for i := range len(pixels[j]) {
			c[j][i] = *model.NewComplex(float64(pixels[j][i]), 0)
		}
	}

	return c
}

func TransposeComplexMatrix(matrix [][]model.Complex) [][]model.Complex {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]model.Complex, cols)
	for i := range transposed {
		transposed[i] = make([]model.Complex, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}
