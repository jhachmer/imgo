package transform

import (
	"math"

	"github.com/jhachmer/imgo/model"
)

// DFT performs a 1D Discrete Fourier Transform on the input slice of complex numbers.
func dft(g []model.Complex, forward bool) []model.Complex {
	M := len(g)
	s := 1 / math.Sqrt(float64(M))

	G := make([]model.Complex, M)

	for m := 0; m < M; m++ {
		sumRe := 0.0
		sumIm := 0.0
		var phim float64 = 2.0 * math.Pi * float64(m) / float64(M)

		for u := 0; u < M; u++ {
			gRe := g[u].Re
			gIm := g[u].Im
			cosw := math.Cos(phim * float64(u))
			sinw := math.Sin(phim * float64(u))
			if !forward {
				sinw = -sinw
			}
			sumRe += gRe*cosw + gIm*sinw
			sumIm += gIm*cosw - gRe*sinw
		}
		G[m] = *model.NewComplex(s*sumRe, s*sumIm)
	}

	return G
}

func DFT2D(g [][]model.Complex, forward bool) [][]model.Complex {
	rows := len(g)
	cols := len(g[0])

	G := make([][]model.Complex, rows)
	for i := range G {
		G[i] = make([]model.Complex, cols)
	}

	for i := 0; i < rows; i++ {
		G[i] = dft(g[i], forward)
	}

	GT := transpose(G)

	for i := 0; i < cols; i++ {
		GT[i] = dft(GT[i], forward)
	}

	G = transpose(GT)

	return G
}

func transpose(matrix [][]model.Complex) [][]model.Complex {
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

func DFTtoReal(c [][]model.Complex) [][]uint8 {
	rows := len(c)
	cols := len(c[0])
	magnitudes := make([][]float64, rows)

	// Compute magnitudes of the complex numbers
	for j := 0; j < rows; j++ {
		magnitudes[j] = make([]float64, cols)
		for i := 0; i < cols; i++ {
			magnitudes[j][i] = c[j][i].Abs()
		}
	}

	// Find the maximum magnitude for normalization
	maxMagnitude := 0.0
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if magnitudes[j][i] > maxMagnitude {
				maxMagnitude = magnitudes[j][i]
			}
		}
	}

	// Normalize to the range [0, 255]
	normalized := make([][]uint8, rows)
	for j := 0; j < rows; j++ {
		normalized[j] = make([]uint8, cols)
		for i := 0; i < cols; i++ {
			normalized[j][i] = uint8((magnitudes[j][i] / maxMagnitude) * 255)
		}
	}

	return normalized
}
