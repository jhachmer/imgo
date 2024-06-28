package types

import (
	"math"
	"math/cmplx"
)

type Gradient2D struct {
	X float64
	Y float64
}

func (g Gradient2D) CalcMagnitude() int {
	v := math.Ceil(math.Hypot(g.X, g.Y))
	if v > 255 {
		v = 255
	}
	return int(v)
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

// ComplexRect returns the complex number x with polar coordinates r, θ as type Complex.
func ComplexRect(mag, pha float64) *Complex {
	val := cmplx.Rect(mag, pha)
	return &Complex{
		Re: real(val),
		Im: imag(val),
	}
}
