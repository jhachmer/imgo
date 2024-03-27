package edge

import (
	"image"

	m "github.com/jhachmer/gocv/model"
)

// SobelOperator Applies Sobel Kernel to given (grayscale) image
// Returns 2D-Slice containing Gradient-(Vectors) for each pixel
func SobelOperator(grayImg *image.Gray) [][]m.Gradient2D {
	var (
		boundsMinX = grayImg.Bounds().Min.X
		boundsMinY = grayImg.Bounds().Min.Y
		boundsMaxX = grayImg.Bounds().Max.X
		boundsMaxY = grayImg.Bounds().Max.Y
	)
	kernelX := [][]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	kernelY := [][]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	K := len(kernelX[0]) / 2
	L := len(kernelX) / 2

	grad2D := make([][]m.Gradient2D, boundsMaxY)
	for i := 0; i < len(grad2D); i++ {
		grad2D[i] = make([]m.Gradient2D, boundsMaxX)
	}

	for v := boundsMinY + L; v < boundsMaxY-L-1; v++ {
		for u := boundsMinX + K; u < boundsMaxX-K-1; u++ {
			var sumGradX int = 0
			var sumGradY int = 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					sourcePix := grayImg.GrayAt(u+i, v+j).Y
					// X-Kernel
					kernX := kernelX[j+L][i+K]
					sumGradX = sumGradX + int(sourcePix)*kernX
					// Y-Kernel
					kernY := kernelY[j+L][i+K]
					sumGradY = sumGradY + int(sourcePix)*kernY
				}
			}
			// Adjust to negative number range
			sumGradX += 127
			sumGradY += 127
			//clamp if necessary
			if sumGradX < 0 {
				sumGradX = 0
			}
			if sumGradX > 255 {
				sumGradX = 255
			}
			if sumGradY < 0 {
				sumGradY = 0
			}
			if sumGradY > 255 {
				sumGradY = 255
			}
			grad2D[v][u].X = sumGradX
			grad2D[v][u].Y = sumGradY
		}
	}

	return grad2D
}

func CalcMagnitudeFromGradient(grad [][]m.Gradient2D) [][]uint8 {
	mag2D := make([][]uint8, len(grad))
	for i := 0; i < len(mag2D); i++ {
		mag2D[i] = make([]uint8, len(grad[i]))
	}

	for v := 0; v < len(grad); v++ {
		for u := 0; u < len(grad[v]); u++ {
			mag2D[v][u] = grad[v][u].CalcMagnitude()
		}
	}

	return mag2D
}
