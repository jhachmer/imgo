package transform

import (
	"math"
	"slices"
)

func HoughTransformLines(pixel [][]uint8, m, n, aMIN int) [][]int {
	M := len(pixel[0])
	N := len(pixel)
	var xR float64 = float64(M) / 2.0
	var yR float64 = float64(N) / 2.0
	var dPHI float64 = math.Pi / float64(m)
	dR := math.Sqrt(float64((M*M)+(N*N))) / float64(n)
	j0 := n / 2

	A := make([][]int, N-1)
	for i := 0; i < N-1; i++ {
		A[i] = make([]int, M-1)
	}

	for u := 0; u < M-1; u++ {
		for v := 0; v < N-1; v++ {
			if pixel[v][u] > 0 {
				x := float64(u) - xR
				y := float64(v) - yR
				for i := 0; i < m-1; i++ {
					phi := dPHI * float64(i)
					r := x*math.Cos(phi) + y*math.Sin(phi)
					j := j0 + int(math.Round(r/dR))
					A[j][i] = A[j][i] + 1
				}
			}
		}
	}

	return A
}

func ScaledAccumulator(A [][]int) [][]uint8 {
	var curMax int = 0
	var N int = len(A)
	var M int = len(A[0])

	for j := 0; j < len(A); j++ {
		subMax := slices.Max(A[j])
		if subMax > curMax {
			curMax = subMax
		}
	}

	scaledA := make([][]uint8, N-1)
	for i := 0; i < N-1; i++ {
		scaledA[i] = make([]uint8, M-1)
	}

	var factor float64 = 255.0 / float64(curMax)

	for v := 0; v < N-1; v++ {
		for u := 0; u < M-1; u++ {
			scaledA[v][u] = uint8(float64(A[v][u]) * factor)

		}
	}

	return scaledA
}
