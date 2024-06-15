package mathutil

import (
	"math"
)

type Gradient2D struct {
	X float64
	Y float64
}

func (g Gradient2D) CalcMagnitude() uint8 {
	v := uint32(math.Ceil(math.Sqrt((g.X * g.X) + (g.Y * g.Y))))

	if v > 250 {
		v = 255
	}
	return uint8(v)
}

type Complex struct {
	Re float64
	Im float64
}

func NewComplex(re float64, im float64) *Complex {
	return &Complex{Re: re, Im: im}
}

func (c Complex) Abs() float64 {
	return math.Hypot(c.Im, c.Re)
}

func (c Complex) Phase() float64 {
	return math.Atan2(c.Im, c.Re)
}
