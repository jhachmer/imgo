package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"slices"

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
func CreateNewGrayFromGradient(gradP [][]m.Gradient2D) (x, y *image.Gray) {
	var (
		bounds = image.Rect(0, 0, len(gradP[0]), len(gradP))
		gradX  = image.NewGray(bounds)
		gradY  = image.NewGray(bounds)
	)

	for posY := 0; posY < bounds.Max.Y; posY++ {
		for posX := 0; posX < bounds.Max.X; posX++ {
			gradX.SetGray(posX, posY, color.Gray{Y: uint8(gradP[posY][posX].X)})
			gradY.SetGray(posX, posY, color.Gray{Y: uint8(gradP[posY][posX].Y)})
		}
	}
	return gradX, gradY
}

func CreateNewMagnitudeFromGradient(pixels [][]uint8) *image.Gray {
	gray := image.NewGray(image.Rect(0, 0, len(pixels[0]), len(pixels)))
	var max uint8 = 0
	for j := range pixels {
		subMax := slices.Max(pixels[j])
		if subMax > max {
			max = subMax
		}
	}
	var factor float64 = 255.0 / float64(max)
	fmt.Println(factor)
	for u := 0; u < gray.Bounds().Max.X; u++ {
		for v := 0; v < gray.Bounds().Max.Y; v++ {
			magValue := uint8(float64(pixels[v][u]) * factor)
			gray.SetGray(u, v, color.Gray{Y: magValue})
		}
	}
	return gray
}
