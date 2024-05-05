package model

import (
	"math"
)

type Gradient2D struct {
	X int
	Y int
}

type Kernel struct {
	size   int
	values [][]int
}

func (g Gradient2D) CalcMagnitude() uint8 {
	var gX int = int(g.X - 127)
	var gY int = int(g.Y - 127)
	fgX := float64(gX) * (1.0 / 8.0)
	fgY := float64(gY) * (1.0 / 8.0)
	det := float64((fgX * fgX) + (fgY * fgY))
	return uint8(math.Sqrt(det))
}
