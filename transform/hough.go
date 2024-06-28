package transform

import (
	"github.com/jhachmer/imgo/internal/ops"
	"math"
	"slices"
	"sync"
)

type HoughTransform struct {
	Accumulator [][]int
}

func NewHoughTransform(input [][]uint8, m, n int) *HoughTransform {
	hough := &HoughTransform{}
	hough.Accumulator = HoughLines(input, m, n)
	return hough
}

func (h *HoughTransform) Output() [][]uint8 {
	return ScaleAccumulator(h.Accumulator)
}

// HoughLines transforms binary image input (2D-slice of uint8's) to hough space
func HoughLines(pixel [][]uint8, m, n int) [][]int {
	M := len(pixel[0])
	N := len(pixel)
	xR := M / 2
	yR := N / 2
	dPHI := math.Pi / float64(m)
	dR := math.Hypot(float64(M), float64(N)) / float64(n)
	j0 := n / 2
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	var wg sync.WaitGroup
	sinCache := make([]float64, m)
	cosCache := make([]float64, m)
	for i := 0; i < m; i++ {
		phi := dPHI * float64(i)
		sinCache[i] = math.Sin(phi)
		cosCache[i] = math.Cos(phi)
	}
	for v := 0; v < N; v++ {
		for u := 0; u < M; u++ {
			if pixel[v][u] > 0 {
				x := xR - u
				y := v - yR
				wg.Add(1)
				go func(x, y int) {
					defer wg.Done()
					for i := 0; i < m; i++ {
						r := float64(x)*cosCache[i] + float64(y)*sinCache[i]
						j := j0 + int(math.Round(r/dR))
						if j >= 0 && j < n {
							A[j][i]++
						}
					}
				}(x, y)
			}
		}
	}
	wg.Wait()
	return A
}

// ScaleAccumulator scales accumulator to uint8 range for image output
func ScaleAccumulator(A [][]int) [][]uint8 {
	var curMax = 0
	N := len(A)
	M := len(A[0])
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < len(A); j++ {
			subMax := slices.Max(A[j])
			if subMax > curMax {
				curMax = subMax
			}
		}
	}()
	scaledA := ops.GenerateSlice[uint8](M, N)
	wg.Wait()
	factor := 255.0 / float64(curMax)
	for v := 0; v < N-1; v++ {
		for u := 0; u < M-1; u++ {
			scaledA[v][u] = uint8(float64(A[v][u]) * factor)
		}
	}
	return scaledA
}
