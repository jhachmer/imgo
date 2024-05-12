package edge

import (
	"errors"
	"image"

	"github.com/jhachmer/gocv/model"
	"github.com/jhachmer/gocv/utils"
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
var SOBEL_COEFF_SUM = 8

// SobelOperator Applies Sobel Kernel to given (grayscale) image
// Returns 2D-Slice containing Gradient-(Vectors) for each pixel
func SobelOperator(grayImg *image.Gray) ([][]model.Gradient2D, error) {
	var (
		boundsMinX = grayImg.Bounds().Min.X
		boundsMinY = grayImg.Bounds().Min.Y
		boundsMaxX = grayImg.Bounds().Max.X
		boundsMaxY = grayImg.Bounds().Max.Y
	)

	kernelX := SOBEL_KERNEL_X
	kernelY := SOBEL_KERNEL_Y

	// Should never happen, but who knwos?
	if kernelX.Size != kernelY.Size {
		return nil, errors.New("x and y kernel sizes dont match")
	}

	K := kernelX.XLen / 2
	L := kernelX.YLen / 2

	// Allocate 2D Array for Gradient Values
	grad2D := make([][]model.Gradient2D, boundsMaxY)
	for i := 0; i < len(grad2D); i++ {
		grad2D[i] = make([]model.Gradient2D, boundsMaxX)
	}
	var (
		sumGradX int
		sumGradY int
	)
	//for v := boundsMinY + L; v < boundsMaxY-L-1; v++ {
	//	for u := boundsMinX + K; u < boundsMaxX-K-1; u++ {
	for v := boundsMinY; v < boundsMaxY; v++ {
		for u := boundsMinX; u < boundsMaxX; u++ {
			sumGradX = 0
			sumGradY = 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					var sourcePix uint8
					if u+i < 0 || v+j < 0 || u+i > grayImg.Bounds().Max.X-1 || v+j > grayImg.Bounds().Max.Y-1 {
						sourcePix = 127
					} else {
						sourcePix = grayImg.GrayAt(u+i, v+j).Y
					}

					// X-Kernel
					kernX := kernelX.Values[j+L][i+K]
					sumGradX = sumGradX + int(sourcePix)*kernX
					// Y-Kernel
					kernY := kernelY.Values[j+L][i+K]
					sumGradY = sumGradY + int(sourcePix)*kernY
				}
			}
			sumGradX /= SOBEL_COEFF_SUM
			sumGradY /= SOBEL_COEFF_SUM
			// Adjust to negative number range
			sumGradX += 127
			sumGradY += 127
			// clamp if necessary
			sumGradX = utils.ClampPixel(sumGradX, 255, 0)
			sumGradY = utils.ClampPixel(sumGradY, 255, 0)
			//fmt.Printf("SumGradX: %v\n", sumGradX)
			//fmt.Printf("SumGradY: %v\n", sumGradY)
			grad2D[v][u].X = sumGradX
			grad2D[v][u].Y = sumGradY
		}
	}

	return grad2D, nil
}

func BuildGradientMagnitudeSlice(grad [][]model.Gradient2D, scale int) [][]uint8 {
	mag2D := make([][]uint8, len(grad))
	for i := 0; i < len(mag2D); i++ {
		mag2D[i] = make([]uint8, len(grad[i]))
	}

	for v := 0; v < len(grad); v++ {
		for u := 0; u < len(grad[v]); u++ {
			mag2D[v][u] = grad[v][u].CalcMagnitude(1)
		}
	}

	return mag2D
}
