package edge

import (
	"math"

	"github.com/jhachmer/imgo/model"
	"github.com/jhachmer/imgo/util"
)

// getOrientationSector returns neighbouring pixel of current pixel in gradient direction
// Returns sector in [0,1,2,3]
// gradient gets rotated and negated if necessary
// Calculates sector with comparisons
func getOrientationSector(dx, dy float64) int {
	dxRot := (math.Cos(math.Pi/8) - math.Sin(math.Pi/8)) * float64(dx)
	dyRot := (math.Sin(math.Pi/8) + math.Cos(math.Pi/8)) * float64(dy)

	if dyRot < 0 {
		dxRot = -dxRot
		dyRot = -dyRot
	}
	var sAng int
	if dxRot >= 0 && dxRot >= dyRot {
		sAng = 0
	}
	if dxRot >= 0 && dxRot < dyRot {
		sAng = 1
	}
	if dxRot < 0 && -dxRot < dyRot {
		sAng = 2
	}
	if dxRot < 0 && -dxRot >= dyRot {
		sAng = 3
	}
	return sAng
}

// isLocalMax returns true if pixel at coordinates (u,v) is a local maximum
// that is the greatest magnitude of the pixel along the gradient direction
func isLocalMax(eMAG [][]uint8, u, v, sAng, tLow int) bool {
	if u <= 0 || u >= len(eMAG[0])-1 || v <= 0 || v >= len(eMAG)-1 {
		return false
	}

	var ml, mr uint8
	mc := eMAG[v][u]
	if int(mc) < tLow {
		return false
	}

	switch sAng {
	case 0:
		ml = eMAG[v][u-1]
		mr = eMAG[v][u+1]
	case 1:
		ml = eMAG[v-1][u-1]
		mr = eMAG[v+1][u+1]
	case 2:
		ml = eMAG[v-1][u]
		mr = eMAG[v+1][u]
	case 3:
		ml = eMAG[v+1][u-1]
		mr = eMAG[v-1][u+1]
	}
	return (ml < mc) && (mc > mr)
}

func traceAndTreshold(eNMS, eBIN [][]uint8, u0, v0, tLow, M, N int) {
	eBIN[v0][u0] = 255
	uL := max(u0-1, 0)
	uR := min(u0+1, M-1)
	vT := max(v0-1, 0)
	vB := min(v0+1, N-1)
	for u := uL; u <= uR; u++ {
		for v := vT; v < vB; v++ {
			if (eNMS[v][u] > uint8(tLow)) && eBIN[v][u] == 0 {
				traceAndTreshold(eNMS, eBIN, u, v, tLow, M, N)
			}
		}
	}
}

func CannyEdgeDetector(grad [][]model.Gradient2D, eMAG [][]uint8, tLOW, tHIGH int) [][]uint8 {
	M := len(grad[0])
	N := len(grad)
	eNMS := make([][]uint8, N)
	for i := 0; i < len(eNMS); i++ {
		eNMS[i] = make([]uint8, M)
	}
	eBIN := make([][]uint8, N)
	for i := 0; i < len(eBIN); i++ {
		eBIN[i] = make([]uint8, M)
	}
	for u := 0; u < M-1; u++ {
		for v := 0; v < N-1; v++ {
			eNMS[v][u] = 0
			eBIN[v][u] = 0
		}
	}

	for u := 1; u < M-2; u++ {
		for v := 1; v < N-2; v++ {
			dX := grad[v][u].X
			dY := grad[v][u].Y
			sO := getOrientationSector(dX, dY)
			if isLocalMax(eMAG, u, v, sO, tLOW) {
				eNMS[v][u] = eMAG[v][u]
			}
		}
	}

	for u := 1; u < M-2; u++ {
		for v := 1; v < N-2; v++ {
			if (eNMS[v][u] >= uint8(tHIGH)) && eBIN[v][u] == 0 {
				traceAndTreshold(eNMS, eBIN, u, v, tLOW, M, N)
			}
		}
	}
	return eBIN
}
