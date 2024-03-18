package filter

import (
	"github.com/jhachmer/gocv/utils"
	"image"
	"image/color"
	"math"
)

// Apply2DFilter Applies given Filter to input image
func Apply2DFilter(path string, filter [][]int, outputFileName string) {
	var filterMatrixSum int
	for _, outer := range filter {
		for _, inner := range outer {
			filterMatrixSum += inner
		}
	}
	s := 1.0 / float64(filterMatrixSum)

	grayImg := utils.ConvertToGrayscaleImage(path)

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
	utils.WriteToFile(outputFileName, newImage)
}
