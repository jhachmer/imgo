package filter

import (
	"image"
	"image/color"
	"math"

	m "github.com/jhachmer/gocv/model"
)

// Apply2DFilter Applies given Filter to input image
func Apply2DFilter(grayImg *image.Gray, filter [][]int) *image.Gray {
	var filterMatrixSum int
	for _, outer := range filter {
		for _, inner := range outer {
			filterMatrixSum += inner
		}
	}
	s := 1.0 / float64(filterMatrixSum)

	newImage := image.NewGray(grayImg.Bounds())

	K := len(filter[0]) / 2
	L := len(filter) / 2

	for v := grayImg.Bounds().Min.Y + L; v < grayImg.Bounds().Max.Y-L; v++ {
		for u := grayImg.Bounds().Min.X + K; u < grayImg.Bounds().Max.X-K; u++ {
			sum := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					p := int(grayImg.GrayAt(u+i, v+j).Y)
					c := filter[j+L][i+K]
					sum = sum + c*p
				}
			}
			q := int(math.Round(s * float64(sum)))
			if q < 0 {
				q = 0
			}
			if q > 255 {
				q = 255
			}
			newImage.SetGray(u, v, color.Gray{Y: uint8(q)})
		}
	}
	return newImage
}

func SobelOperator(grayImg *image.Gray) *[][]m.Gradient2D {
	kernelX := [][]int{
		{-1, 0, 1},
		{-1, 0, 1},
		{-1, 0, 1},
	}
	kernelY := [][]int{
		{-1, -1, -1},
		{0, 0, 0},
		{1, 1, 1},
	}

	K := len(kernelX[0]) / 2
	L := len(kernelX) / 2

	grad2D := make([][]m.Gradient2D, grayImg.Bounds().Max.Y)
	for i := range grad2D {
		grad2D[i] = make([]m.Gradient2D, grayImg.Bounds().Max.X)
	}

	for v := grayImg.Bounds().Min.Y + L; v < grayImg.Bounds().Max.Y-L; v++ {
		for u := grayImg.Bounds().Min.X + K; u < grayImg.Bounds().Max.X-K; u++ {
			sumGradX := 0
			sumGradY := 0
			for j := -L; j <= L; j++ {
				for i := -K; i <= K; i++ {
					sourcePix := int(grayImg.GrayAt(u+i, v+j).Y)
					// X-Kernel
					kernX := kernelX[j+L][i+K]
					sumGradX = sumGradX + sourcePix*kernX
					// Y-Kernel
					kernY := kernelY[j+L][i+K]
					sumGradY = sumGradY + sourcePix*kernY
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
			grad2D[u][v].X = uint8(sumGradX)
			grad2D[u][v].Y = uint8(sumGradY)
		}
	}

	return &grad2D
}
