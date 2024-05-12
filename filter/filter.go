package filter

import (
	"image"
	"image/color"
	"math"

	"github.com/jhachmer/gocv/model"
	"github.com/jhachmer/gocv/utils"
)

// Apply2DFilter Applies given Filter to input (grayscale) image.
// Returns grayscale image with applied filter
func Apply2DFilter(grayImg *image.Gray, k model.Kernel2D) *image.Gray {
	// Scale by sum of filter coefficients
	s := 1.0 / float64(k.Size)

	newImage := image.NewGray(grayImg.Bounds())

	K := k.XLen / 2
	L := k.YLen / 2

	//for v := grayImg.Bounds().Min.Y + L; v < grayImg.Bounds().Max.Y-L; v++ {
	//	for u := grayImg.Bounds().Min.X + K; u < grayImg.Bounds().Max.X-K; u++ {
	for v := grayImg.Bounds().Min.Y; v < grayImg.Bounds().Max.Y; v++ {
		for u := grayImg.Bounds().Min.X; u < grayImg.Bounds().Max.X; u++ {
			sum := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					var p int
					if u+i < 0 || v+j < 0 || u+i > grayImg.Bounds().Max.X-1 || v+j > grayImg.Bounds().Max.Y-1 {
						p = 127
					} else {
						p = int(grayImg.GrayAt(u+i, v+j).Y)
					}
					c := k.Values[j+L][i+K]
					sum = sum + c*p
				}
			}
			// Clamping if necessary
			q := int(math.Round(s * float64(sum)))
			q = utils.ClampPixel(q, 255, 0)
			newImage.SetGray(u, v, color.Gray{Y: uint8(q)})
		}
	}
	return newImage
}
