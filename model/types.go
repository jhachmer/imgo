package model

import (
	"math"
)

type Gradient2D struct {
	X float64
	Y float64
}

type Imaginary struct {
	Re float64
	Im float64
}

type Kernel2D struct {
	Size   int
	Values [][]int
	XLen   int
	YLen   int
}

type Kernel1D struct {
	Size   int
	Values []int
	Len    int
}

func (g Gradient2D) CalcMagnitude() uint8 {
	// Adjust number range

	v := uint32(math.Ceil(math.Sqrt(float64((g.X * g.X) + (g.Y * g.Y)))))

	if v > 250 {
		v = 255
	} else if v < 5 {
		v = 0
	}
	return uint8(v)
}

func (k Kernel2D) CalcCoeffSum() int {
	var sum int
	for i := range k.Values {
		for j := range k.Values[i] {
			sum += k.Values[i][j]
		}
	}

	return sum
}

func (k Kernel1D) CalcCoeffSum() int {
	var sum int
	for i := range k.Values {
		sum += k.Values[i]
	}

	return sum
}

func NewKernel2D(values [][]int) *Kernel2D {
	k := new(Kernel2D)
	k.Values = values
	k.Size = k.CalcCoeffSum()
	k.XLen = len(values[0])
	k.YLen = len(values)

	return k
}

func NewKernel1D(values []int) *Kernel1D {
	k := new(Kernel1D)
	k.Values = values
	k.Size = k.CalcCoeffSum()
	k.Len = len(values)

	return k
}
