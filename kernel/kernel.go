package kernel

import (
	"errors"

	"github.com/jhachmer/imgo/mathutil"
)

type Kernel2D struct {
	Size   int
	Values [][]int
	Width  int
	Height int
}

func NewKernel2D(values [][]int) (*Kernel2D, error) {
	k := new(Kernel2D)
	k.Values = values
	k.CalcCoefficientSum()
	k.Width = len(values[0])
	k.Height = len(values)
	if k.Width != k.Height {
		return nil, errors.New("dimensions of kernel are not symmetrical")
	}

	return k, nil
}

func (k *Kernel2D) CalcCoefficientSum() {
	var sum int
	for i := range k.Values {
		for j := range k.Values[i] {
			//sum += mathutil.Abs(k.Values[i][j])
			sum += k.Values[i][j]

		}
	}
	k.Size = sum
}

func (k *Kernel2D) GetHalfKernelSize() (int, int) {
	K := k.Width / 2
	L := k.Height / 2
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
	err := k.CalcCoefficientSum()
	if err != nil {
		return nil, err
	}
	k.Len = len(values)

	return k, nil
}

func (k *Kernel1D) GetHalfKernelSize() int {
	return k.Len / 2

}

func (k *Kernel1D) CalcCoefficientSum() error {
	var sum int
	for i := range k.Values {
		sum += mathutil.Abs(k.Values[i])
	}
	if sum == 0 {
		return errors.New("sum of filter coefficients is zero")
	}

	k.Size = sum

	return nil
}
