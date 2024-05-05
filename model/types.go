package model

import (
	"math"
)

type Gradient2D struct {
	X int
	Y int
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

func (g Gradient2D) CalcMagnitude() uint8 {
	var gX int = int(g.X - 127)
	var gY int = int(g.Y - 127)
	fgX := float64(gX) * (1.0 / 8.0)
	fgY := float64(gY) * (1.0 / 8.0)
	det := float64((fgX * fgX) + (fgY * fgY))
	return uint8(math.Sqrt(det))
}

func (k Kernel2D) CalcCoeffSum() float32 {
	var sum int
	for i := range k.Values {
		for j := range k.Values[i] {
			sum += k.Values[i][j]
		}
	}
	return 1. / float32(sum)
}

func (k Kernel1D) CalcCoeffSum() float32 {
	var sum int
	for i := range k.Values {
		sum += k.Values[i]
	}
	return 1. / float32(sum)
}
