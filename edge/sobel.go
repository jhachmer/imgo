package edge

import (
	"errors"
	"image"
	"math"

	"github.com/jhachmer/imgo/model"
	"github.com/jhachmer/imgo/utils"
)

func sobelKernelXAndY() (kernX, kernY *model.Kernel2D) {

    sobelKernelX, err := model.NewKernel2D([][]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
})
    if err != nil {
        panic(err)
    }

    sobelKernelY, err := model.NewKernel2D([][]int{
	{-1, -2, -1},
	{0, 0, 0},
	{1, 2, 1},
})
    
    if err != nil {
        panic("invalid kernel")
    }
	return sobelKernelX, sobelKernelY
}

// SobelOperator Applies Sobel Kernel to given (grayscale) image
// Returns 2D-Slice containing Gradient-(Vectors) for each pixel
func SobelOperator(grayImg *image.Gray) ([][]model.Gradient2D, error) {
	var (
		boundsMaxX = grayImg.Bounds().Max.X
		boundsMaxY = grayImg.Bounds().Max.Y
		K          int
		L          int
    )

	kernelX, kernelY := sobelKernelXAndY()

	// Should never happen, but who knows?
	if kernelX.Size != kernelY.Size {
		return nil, errors.New("x and y kernel sizes dont match")
	}

	xK, xL := kernelX.GetHalfKernelSize()
    yK, yL := kernelY.GetHalfKernelSize()
    if xK != yK && xL != yL {
		return nil, errors.New("x- and y-Kernel dimensions do not match")
    } else {
        K, L = xK, xL
    }

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
					if u+i < 0 || v+j < 0 || u+i >= boundsMaxX || v+j >= boundsMaxY {
						xPix, yPix = utils.BorderDetection(u, v, i, j, boundsMaxX, boundsMaxY)
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
			uFX, uFY := math.Abs(float64(sumGradX)), math.Abs(float64(sumGradY))

			// uFX, uFY = uFX/float64(SOBEL_COEFF_SUM), uFY/float64(SOBEL_COEFF_SUM)
			grad2D[v][u].X = uFX
			grad2D[v][u].Y = uFY
		}
	}

	return grad2D, nil
}

func BuildGradientMagnitudeSlice(grad [][]model.Gradient2D) [][]uint8 {
	mag2D := utils.GeneratePixelSlice(len(grad[0]), len(grad))

	var wg sync.WaitGroup

	for v := 0; v < len(grad); v++ {
		for u := 0; u < len(grad[v]); u++ {
			mag2D[v][u] = grad[v][u].CalcMagnitude()
		}
	}

	return mag2D
}
