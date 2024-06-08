package filter

import (
	"image"
	"image/color"
	"math"

	"github.com/jhachmer/imgo/model"
	"github.com/jhachmer/imgo/utils"
)

// Apply2DFilterToGray applies given 2D-Filter to input (grayscale) image.
// Returns grayscale image with applied filter
func Apply2DFilterToGray(grayImg *image.Gray, k *model.Kernel2D) *image.Gray {
	var (
		s          = 1.0 / float64(k.Size)
		boundsMaxX = grayImg.Bounds().Max.X
		boundsMaxY = grayImg.Bounds().Max.Y
		newImage   = image.NewGray(grayImg.Bounds())
		K, L       = k.GetHalfKernelSize()
		xPix       int
		yPix       int
	)

	for v := 0; v < boundsMaxY; v++ {
		for u := 0; u < boundsMaxX; u++ {
			sum := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					if u+i < 0 || v+j < 0 || u+i >= boundsMaxX || v+j >= boundsMaxY {
						xPix, yPix = utils.BorderDetection(u, v, i, j, boundsMaxX, boundsMaxY)
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
			// Clamping if necessary
			q = utils.ClampPixel(q, 255, 0)
			newImage.SetGray(u, v, color.Gray{Y: uint8(q)})
		}
	}
	return newImage
}
