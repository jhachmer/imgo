package edge

import (
	"errors"
	"image"

	"github.com/jhachmer/gocvlite/model"
	"github.com/jhachmer/gocvlite/utils"
)

var SOBEL_KERNEL_X model.Kernel2D = *model.NewKernel2D([][]int{
	{-3, 0, 3},
	{-10, 0, 10},
	{-3, 0, 3},
})

var SOBEL_KERNEL_Y model.Kernel2D = *model.NewKernel2D([][]int{
	{-3, -10, -3},
	{0, 0, 0},
	{3, 10, 3},
})

var SOBEL_HALF_KERNEL_SIZE = SOBEL_KERNEL_X.XLen / 2
var SOBEL_COEFF_SUM = 32

// SobelOperator Applies Sobel Kernel to given (grayscale) image
// Returns 2D-Slice containing Gradient-(Vectors) for each pixel
func SobelOperator(grayImg *image.Gray) ([][]model.Gradient2D, error) {
	var (
		boundsMaxX = grayImg.Bounds().Max.X
		boundsMaxY = grayImg.Bounds().Max.Y
	)

	kernelX := SOBEL_KERNEL_X
	kernelY := SOBEL_KERNEL_Y

	// Should never happen, but who knows?
	if kernelX.Size != kernelY.Size {
		return nil, errors.New("x and y kernel sizes dont match")
	}

	K := SOBEL_HALF_KERNEL_SIZE //kernelX.XLen / 2
	L := SOBEL_HALF_KERNEL_SIZE //kernelX.YLen / 2

	// Allocate 2D Slice for Gradient Values
	grad2D := make([][]model.Gradient2D, boundsMaxY)
	for i := 0; i < len(grad2D); i++ {
		grad2D[i] = make([]model.Gradient2D, boundsMaxX)
	}
	var (
		sumGradX int
		sumGradY int
		xPix     int
		yPix     int
	)

	for v := 0; v < boundsMaxY; v++ {
		for u := 0; u < boundsMaxX; u++ {
			sumGradX = 0
			sumGradY = 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					if u+i < 0 || v+j < 0 || u+i >= boundsMaxX-1 || v+j >= boundsMaxY-1 {
						xPix, yPix = utils.BorderDetection(u, v, i, j, boundsMaxX-1, boundsMaxY-1)
					} else {
						xPix, yPix = u+i, v+j
					}
					sourcePix := grayImg.GrayAt(xPix, yPix).Y

					// X-Kernel
					kernX := kernelX.Values[j+L][i+K]
					sumGradX += int(sourcePix) * kernX
					// Y-Kernel
					kernY := kernelY.Values[j+L][i+K]
					sumGradY += int(sourcePix) * kernY
				}
			}
			//sumGradX /= SOBEL_COEFF_SUM
			//sumGradY /= SOBEL_COEFF_SUM
			// Adjust to negative number range
			//sumGradX += 127
			//sumGradY += 127
			// clamp if necessary
			//sumGradX = utils.ClampPixel(sumGradX, 255, 0)
			//sumGradY = utils.ClampPixel(sumGradY, 255, 0)
			grad2D[v][u].X = sumGradX
			grad2D[v][u].Y = sumGradY
		}
	}

	return grad2D, nil
}

func BuildGradientMagnitudeSlice(grad [][]model.Gradient2D) [][]uint8 {
	mag2D := make([][]uint8, len(grad))
	for i := 0; i < len(mag2D); i++ {
		mag2D[i] = make([]uint8, len(grad[i]))
	}
	for v := 0; v < len(grad); v++ {
		for u := 0; u < len(grad[v]); u++ {
			mag2D[v][u] = grad[v][u].CalcMagnitude(SOBEL_COEFF_SUM)
		}
	}

	return mag2D
}
