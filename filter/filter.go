package filter

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/jhachmer/gocvlite/model"
	"github.com/jhachmer/gocvlite/utils"
)

// Apply2DFilter Applies given Filter to input (grayscale) image.
// Returns grayscale image with applied filter
func Apply2DFilter(grayImg *image.Gray, k model.Kernel2D) *image.Gray {
	// Scale by sum of filter coefficients
	s := 1.0 / float64(k.Size)

	newImage := image.NewGray(grayImg.Bounds())

	K := k.XLen / 2
	L := k.YLen / 2
	fmt.Println(K)
	fmt.Println(L)
	var (
		xPix int
		yPix int
	)

	for v := 0; v < grayImg.Bounds().Max.Y; v++ {
		for u := 0; u < grayImg.Bounds().Max.X; u++ {
			sum := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					if u+i < 0 || v+j < 0 || u+i > grayImg.Bounds().Max.X || v+j > grayImg.Bounds().Max.Y {
						xPix, yPix = utils.BorderDetection(u, v, i, j, grayImg.Bounds().Max.X, grayImg.Bounds().Max.Y)
					} else {
						xPix, yPix = u+i, v+j
					}
					p := int(grayImg.GrayAt(xPix, yPix).Y)

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
