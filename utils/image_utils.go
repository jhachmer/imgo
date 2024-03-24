package utils

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	m "github.com/jhachmer/gocv/model"
)

func RgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			rgba := img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

func ConvertImageToGrayPNG(path string) *image.Gray {
	sourceFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	sourceImg, err := png.Decode(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	return RgbaToGray(sourceImg)
}

func WriteGrayToFilePNG(outputFileName string, newImage *image.Gray) {
	f, err := os.Create(outputFileName + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, newImage); err != nil {
		log.Fatal(err)
	}
}

// TODO: input type ???? make usable for more
func CreateNewGrayFromGradient(gradP *[][]m.Gradient2D) (x, y *image.Gray) {
	grad := *gradP
	var (
		bounds = image.Rect(0, 0, len(grad[0]), len(grad))
		gradX  = image.NewGray(bounds)
		gradY  = image.NewGray(bounds)
	)
	for posY := range grad {
		for posX := range grad[posY] {
			gradX.SetGray(posX, posY, color.Gray{Y: grad[posX][posY].X})
			gradY.SetGray(posX, posY, color.Gray{Y: grad[posX][posY].Y})
		}
	}
	return gradX, gradY
}
