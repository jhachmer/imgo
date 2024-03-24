package model

import (
	"math"
)

type Gradient2D struct {
	X int
	Y int
}

func (g Gradient2D) CalcMagnitude() uint8 {
	var gX int = int(g.X - 127)
	var gY int = int(g.Y - 127)
	det := float64((gX * gX) + (gY * gY))
	return uint8(math.Sqrt(det))
}
