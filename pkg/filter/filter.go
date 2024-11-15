package filter

import (
	"math"

	"github.com/jhachmer/imgo/internal/ops"
	"github.com/jhachmer/imgo/pkg/border"
	"github.com/jhachmer/imgo/pkg/img"

	"github.com/jhachmer/imgo/pkg/kernel"
)

type Filter struct {
	Pixels [][]uint8
}

func NewFilter(img *img.ImageGray, k *kernel.Kernel2D, derivative bool) *Filter {
	return &Filter{
		Pixels: Apply2DFilter(img.Pixels, k, derivative),
	}
}

var _ img.Outputable = (*Filter)(nil)

func (f Filter) Output() [][]uint8 {
	return f.Pixels
}

// Apply2DFilter applies given 2D-Filter to input (grayscale) image.
// Returns grayscale image with applied filter
func Apply2DFilter(in [][]uint8, k *kernel.Kernel2D, derivative bool) [][]uint8 {
	var (
		s float64

		cols, rows = len(in[0]), len(in)
		out        = ops.GenerateSlice[uint8](cols, rows)
		K, L       = k.GetHalfKernelSize()
		xPix       int
		yPix       int
	)
	if k.Size == 0 {
		s = 1
	} else {
		s = 1.0 / float64(k.Size)
	}

	for v := 0; v <= rows-1; v++ {
		for u := 0; u <= cols-1; u++ {
			sum := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					xPix, yPix = border.Detection(u, v, i, j, cols-1, rows-1)
					p := int(in[yPix][xPix])
					c := k.Values[j+L][i+K]
					sum = sum + c*p
				}
			}
			// Scale by sum of filter coefficients
			q := int(math.Round(s * float64(sum)))
			// adjust to number range if using derivative kernel
			if derivative {
				q += 127
			}
			// Clamping if necessary
			out[v][u] = ops.ClampPixel(q)
			//fmt.Printf("%d ", u)
			//fmt.Printf("%d \n", v)
		}
	}

	return out
}

func CreateGaussKernel1D(sigma float64) *kernel.Kernel1D {
	center := int(3.0 * sigma)
	h := make([]int, 2*center+1)
	sigma2 := sigma * sigma
	for i := range h {
		r := float64(center - i)
		h[i] = int(math.Exp(-0.5*(r*r)/sigma2) * 100)
	}
	k, err := kernel.NewKernel1D(h)
	if err != nil {
		panic(err)
	}
	return k
}
