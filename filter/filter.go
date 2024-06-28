package filter

import (
	"github.com/jhachmer/imgo/border"
	"github.com/jhachmer/imgo/img"
	"github.com/jhachmer/imgo/internal/ops"
	"math"

	"github.com/jhachmer/imgo/kernel"
)

// Apply2DFilterToGray applies given 2D-Filter to input (grayscale) image.
// Returns grayscale image with applied filter
func Apply2DFilterToGray(grayImg *image.Gray, k *kernel.Kernel2D, derivative bool) *image.Gray {
	var (
		s float64

		boundsMaxX = grayImg.Bounds().Max.X
		boundsMaxY = grayImg.Bounds().Max.Y
		newImage   = image.NewGray(grayImg.Bounds())
		K, L       = k.GetHalfKernelSize()
		xPix       int
		yPix       int
	)
	if k.Size == 0 {
		s = 1
	} else {
		s = 1.0 / float64(k.Size)
	}

	for v := 0; v < boundsMaxY; v++ {
		for u := 0; u < boundsMaxX; u++ {
			sum := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					if u+i < 0 || v+j < 0 || u+i >= boundsMaxX || v+j >= boundsMaxY {
						xPix, yPix = border.Detection(u, v, i, j, boundsMaxX, boundsMaxY)
					} else {
						xPix, yPix = u+i, v+j
					}
					p := int(grayImg.GrayAt(xPix, yPix).Y)

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
			q = ops.ClampPixel(q, 255, 0)

			newImage.SetGray(u, v, color.Gray{Y: uint8(q)})
		}
	}
	return newImage
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
