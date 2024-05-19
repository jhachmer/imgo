package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"slices"
)

func ConvertToGrayScale(img image.Image) *image.Gray {
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

func ReadImageFromPath(path string) *image.Image {
	sourceFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	sourceImg, _, err := image.Decode(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	return &sourceImg
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

func CreateMagnitudeGrayImageFromGradient(pixels [][]uint8) *image.Gray {
	gray := image.NewGray(image.Rect(0, 0, len(pixels[0]), len(pixels)))

	var curMax uint8 = 0

	for j := 0; j <= gray.Bounds().Max.Y-1; j++ {
		subMax := slices.Max(pixels[j])
		if subMax > curMax {
			curMax = subMax
		}
	}

	var factor float64 = 255.0 / float64(curMax)
	fmt.Printf("Scale Factor %v\n", factor)
	for u := 0; u < gray.Bounds().Max.X; u++ {
		for v := 0; v < gray.Bounds().Max.Y; v++ {
			magValue := uint8(float64(pixels[v][u]) * factor)
			gray.SetGray(u, v, color.Gray{Y: magValue})
		}
	}
	return gray
}

func SliceToImage(pixels [][]uint8) *image.Gray {
	var (
		bounds = image.Rect(0, 0, len(pixels[0]), len(pixels))
		canny  = image.NewGray(bounds)
	)

	for posY := 0; posY < bounds.Max.Y; posY++ {
		for posX := 0; posX < bounds.Max.X; posX++ {
			canny.SetGray(posX, posY, color.Gray{Y: uint8(pixels[posY][posX])})
		}
	}
	return canny
}
