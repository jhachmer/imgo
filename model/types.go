package model

import (
	"errors"
	"math"

	mathutil "github.com/jhachmer/gocvlite/mathutil"
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

func (k Kernel2D) CalcCoeffSum() (int, error) {
	var sum int
	for i := range k.Values {
		for j := range k.Values[i] {
			sum += mathutil.Abs((k.Values[i][j]))
		}
	}
	if sum == 0 {
		return 0, errors.New("Sum of filter coefficients is zero.")
	}

	return sum, nil
}

func (k Kernel1D) CalcCoeffSum() (int, error) {
	var sum int
	for i := range k.Values {
		sum += mathutil.Abs(k.Values[i])
	}
	if sum == 0 {
		return 0, errors.New("Sum of filter coefficients is zero.")
	}

	return sum, nil
}

func (k Kernel2D) GetHalfKernelSize() (int, int) {
	K := k.XLen / 2
	L := k.YLen / 2
	return K, L
}

func NewKernel2D(values [][]int) (*Kernel2D, error) {
	k := new(Kernel2D)
	k.Values = values
	size, err := k.CalcCoeffSum()
	if err != nil {
		return nil, err
	}
	k.Size = size
	k.XLen = len(values[0])
	k.YLen = len(values)
	if k.XLen != k.YLen {
		return nil, errors.New("Dimensions of Kernel are not symmetrical")
	}

	return k, nil
}

func NewKernel1D(values []int) (*Kernel1D, error) {
	k := new(Kernel1D)
	k.Values = values
	size, err := k.CalcCoeffSum()
	if err != nil {
		return nil, err
	}
	k.Size = size
	k.Len = len(values)

	return k, nil
}
