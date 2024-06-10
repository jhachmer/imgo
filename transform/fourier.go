package transform

import (
	"math"

	"github.com/jhachmer/imgo/model"
	"github.com/jhachmer/imgo/util"
)

// DFT performs a 1D Discrete Fourier Transform on the input slice of complex numbers.
func dft1D(g []model.Complex, forward bool) []model.Complex {
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

	for i := 0; i < rows; i++ {
		g[i] = dft1D(g[i], forward)
	}

	g = util.TransposeComplexMatrix(g)

	for i := 0; i < cols; i++ {
		g[i] = dft1D(g[i], forward)
	}

	g = util.TransposeComplexMatrix(g)

	return g
}

func DFTMagnitude(c [][]model.Complex) [][]float64 {
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

	return magnitudes
}

func OutputTransformation(magnitudes [][]float64, logarithmic bool) [][]uint8 {
	cols, rows := len(magnitudes[0]), len(magnitudes)
	// Find the maximum magnitude for normalization
	maxMagnitude := 0.0
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if magnitudes[j][i] > maxMagnitude {
				maxMagnitude = magnitudes[j][i]
			}
		}
	}

	c := 255 / math.Log(1+math.Abs(maxMagnitude))

	// Normalize to the range [0, 255]
	normalized := make([][]uint8, rows)
	for j := 0; j < rows; j++ {
		normalized[j] = make([]uint8, cols)
		for i := 0; i < cols; i++ {
			if !logarithmic {
				v := int(magnitudes[j][i] / maxMagnitude * 255)
				v = util.ClampPixel(v, 255, 0)
				normalized[j][i] = uint8(v)
			} else {
				v := int(c * math.Log(1+math.Abs(magnitudes[j][i])))
				v = util.ClampPixel(v, 255, 0)
				normalized[j][i] = uint8(v)
			}
		}
	}

	normalized = dftshift(normalized)

	return normalized
}

func dftshift(matrix [][]uint8) [][]uint8 {
	cols, rows := len(matrix[0]), len(matrix)
	shifted := util.GeneratePixelSlice(cols, rows)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			newI := (i + cols/2) % cols
			newJ := (j + rows/2) % rows
			shifted[newJ][newI] = matrix[j][i]
		}
	}
	return shifted
}
