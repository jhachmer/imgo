package ops

import (
	m "github.com/jhachmer/imgo/mathutil"
	"slices"
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

func GenerateComplexSlice(pixels [][]uint8) [][]m.Complex {
	c := make([][]m.Complex, len(pixels))
	for i := range c {
		c[i] = make([]m.Complex, len(pixels[i]))
	}
	for j := range len(pixels) {
		for i := range len(pixels[j]) {
			c[j][i] = *m.NewComplex(float64(pixels[j][i]), 0)
		}
	}

	return c
}

func TransposeComplexMatrix(matrix [][]m.Complex) [][]m.Complex {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]m.Complex, cols)
	for i := range transposed {
		transposed[i] = make([]m.Complex, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

func FindMaxIn2DSlice[T Number](s [][]T) T {
	var subMax, curMax T
	for i := range s {
		subMax = slices.Max(s[i])
		if subMax > curMax {
			curMax = subMax
		}
	}
	return curMax
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}
