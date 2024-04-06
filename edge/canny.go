package edge

import "math"

func getOrientationSector(dx int, dy int) int {
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
