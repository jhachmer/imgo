package model

import (
	"errors"
	"github.com/jhachmer/imgo/mathutil"
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

type Imaginary struct {
	Re float64
	Im float64
}

func NewImaginary(re float64, im float64) *Imaginary {
	return &Imaginary{Re: re, Im: im}
}

type Kernel2D struct {
	Size   int
	Values [][]int
	XLen   int
	YLen   int
}

func NewKernel2D(values [][]int) (*Kernel2D, error) {
	k := new(Kernel2D)
	k.Values = values
	size, err := k.CalcCoefficientSum()
	if err != nil {
		return nil, err
	}
	k.Size = size
	k.XLen = len(values[0])
	k.YLen = len(values)
	if k.XLen != k.YLen {
		return nil, errors.New("dimensions of kernel are not symmetrical")
	}

	return k, nil
}

func (k Kernel2D) CalcCoefficientSum() (int, error) {
	var sum int
	for i := range k.Values {
		for j := range k.Values[i] {
			sum += mathutil.Abs(k.Values[i][j])
		}
	}
	if sum == 0 {
		return 0, errors.New("sum of filter coefficients is zero")
	}

	return sum, nil
}

func (k Kernel2D) GetHalfKernelSize() (int, int) {
	K := k.XLen / 2
	L := k.YLen / 2
	return K, L
}

type Kernel1D struct {
	Size   int
	Values []int
	Len    int
}

func NewKernel1D(values []int) (*Kernel1D, error) {
	k := new(Kernel1D)
	k.Values = values
	size, err := k.CalcCoefficientSum()
	if err != nil {
		return nil, err
	}
	k.Size = size
	k.Len = len(values)

	return k, nil
}

func (k Kernel1D) GetHalfKernelSize() int {
	K := k.Len / 2
	return K
}

func (k Kernel1D) CalcCoefficientSum() (int, error) {
	var sum int
	for i := range k.Values {
		sum += mathutil.Abs(k.Values[i])
	}
	if sum == 0 {
		return 0, errors.New("sum of filter coefficients is zero")
	}

	return sum, nil
}
