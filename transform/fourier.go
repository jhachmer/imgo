package transform

import (
	"math"

	"github.com/jhachmer/imgo/model"
	"github.com/jhachmer/imgo/util"
)

// dft1D performs a 1D Discrete Fourier Transform on the input slice of complex numbers.
// forward flag sets whether or not to use inverse DFT
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

// DFT2D applies DFT to 2D-slice of complex numbers
// real number image slices can be coverted to complex slices using GenerateComplexSlice in util package
// forward flag sets whether or not to use inverse DFT
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

// GenerateComplexSlice calculates the magnitude of every complex number in given 2D-slice
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

// OutputTransformation adjusts numbers in given 2D-slice to uint8-range
// logarithmic flags controls if linear or logarithmic transformation is used
func OutputTransformation(magnitudes [][]float64, logarithmic bool) [][]uint8 {
	cols, rows := len(magnitudes[0]), len(magnitudes)
	var c float64
	maxMagnitude := 0.0
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if magnitudes[j][i] > maxMagnitude {
				maxMagnitude = magnitudes[j][i]
			}
		}
	}

	if logarithmic {
		c = 255 / math.Log(1+math.Abs(maxMagnitude))
	}

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

// dftshift uses symmetry of DFT to allign low-frequency parts of signal (DC) to the center of the image
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
