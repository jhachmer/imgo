package edge

import "math"

type Gradient2D struct {
	X int8
	Y int8
}

type Magnitude int8

func (g *Gradient2D) CalcMagnitude() int8 {
	det := float64((g.X * g.X) + (g.Y * g.Y))
	return int8(math.Sqrt(det))
}
