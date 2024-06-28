package edge

import (
	"slices"
	"sync"

	"github.com/jhachmer/imgo/border"
	"github.com/jhachmer/imgo/internal/ops"
	m "github.com/jhachmer/imgo/internal/types"
	"github.com/jhachmer/imgo/kernel"
)

type Sobel struct {
	Gradient   [][]m.Gradient2D
	Magnitudes [][]uint8
}

func NewSobel(in [][]uint8) *Sobel {
	s := &Sobel{
		Gradient: sobelOperator(in),
	}
	s.Magnitudes = s.SobelMagnitudes()
	return s
}

func (s *Sobel) Output() [][]uint8 {
	return s.Magnitudes
}

func (s *Sobel) SobelMagnitudes() [][]uint8 {
	rows := len(s.Gradient)
	cols := len(s.Gradient[0])

	mag2D := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mag2D[i] = make([]int, cols)
	}

	var wg sync.WaitGroup
	for v := 0; v < len(s.Gradient); v++ {
		for u := 0; u < len(s.Gradient[v]); u++ {
			wg.Add(1)
			go func(v, u int) {
				defer wg.Done()
				mag2D[v][u] = s.Gradient[v][u].CalcMagnitude()
			}(v, u)
		}
	}
	wg.Wait()

	res := ops.GenerateSlice[uint8](cols, rows)
	curMax := 0
	for j := 0; j < len(mag2D); j++ { // Fixed loop condition
		subMax := slices.Max(mag2D[j])
		if subMax > curMax {
			curMax = subMax
		}
	}
	var factor = 255.0 / float64(curMax)
	for u := 0; u < cols; u++ {
		for v := 0; v < rows; v++ {
			res[v][u] = uint8(float64(mag2D[v][u]) * factor)
		}
	}
	return res
}

// SobelOperator applies Sobel Kernel to the given (grayscale) image
// Returns a 2D slice containing Gradient vectors for each pixel
func sobelOperator(input [][]uint8) [][]m.Gradient2D {
	kernelX, err := kernel.NewKernel2D([][]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	})
	if err != nil {
		panic(err)
	}

	kernelY, err := kernel.NewKernel2D([][]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	})
	if err != nil {
		panic(err)
	}

	var (
		cols = len(input[0])
		rows = len(input)
		K, L = kernelX.GetHalfKernelSize()
	)

	// Allocate 2D Slice for Gradient Values
	grad2D := make([][]m.Gradient2D, rows)
	for i := 0; i < len(grad2D); i++ {
		grad2D[i] = make([]m.Gradient2D, cols)
	}

	var (
		sumGradX int
		sumGradY int
	)

	for v := 1; v <= rows-1; v++ {
		for u := 1; u <= cols-1; u++ {
			sumGradX = 0
			sumGradY = 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					xPix, yPix := border.Detection(u, v, i, j, cols-2, rows-2)
					sourcePix := input[yPix][xPix]

					// X-Kernel
					kernX := kernelX.Values[j+L][i+K]
					sumGradX += int(sourcePix) * kernX
					// Y-Kernel
					kernY := kernelY.Values[j+L][i+K]
					sumGradY += int(sourcePix) * kernY
				}
			}
			uFX, uFY := float64(sumGradX), float64(sumGradY)

			grad2D[v][u].X = uFX
			grad2D[v][u].Y = uFY
		}
	}

	return grad2D
}
