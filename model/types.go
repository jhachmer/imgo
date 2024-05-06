package model

import (
	"math"
)

type Gradient2D struct {
	X int
	Y int
}

type Imaginary struct {
	Re float64
	Im float64
}

type Kernel2D struct {
	Size   int
	Values [][]int
}

type Kernel1D struct {
	Size   int
	Values []int
}

type Filter1D struct {
	XKernel Kernel1D
	YKernel Kernel1D
}

type Filter2D struct {
	Kernel Kernel2D
}

func (g Gradient2D) CalcMagnitude(scale int) uint8 {
	var gX int = int(g.X - 127)
	var gY int = int(g.Y - 127)
	fgX := float64(gX) * (1.0 / float64(scale))
	fgY := float64(gY) * (1.0 / float64(scale))
	det := float64((fgX * fgX) + (fgY * fgY))
	return uint8(math.Sqrt(det))
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
	return k
}

func NewKernel1D(values []int) *Kernel1D {
	k := new(Kernel1D)
	k.Values = values
	k.Size = k.CalcCoeffSum()
	return k
}
