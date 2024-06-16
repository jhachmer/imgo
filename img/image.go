package img

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"slices"

	"github.com/jhachmer/imgo/ops"
)

type Outputer interface {
	Output() [][]uint8
}

// ConvertToGrayScale Converts input to grayscale image
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

// ReadImageFromPath reads an image from filesystem
// Requires filepath from working directory
// Returns image as pointer to Go-type image.Image
func ReadImageFromPath(path string) image.Image {
	sourceFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	sourceImg, _, err := image.Decode(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	return sourceImg
}

// ToPNG  writes Go image to filesystem
// Creates image at given filepath
func ToPNG(outputFileName string, newImage *image.Gray) {
	f, err := os.Create(outputFileName + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, newImage); err != nil {
		log.Fatal(err)
	}
}

// ToImage  converts a 2D-Slice to an image
// can be saved to file system using decoder
func ToImage(output Outputer) *image.Gray {
	var (
		pixels = output.Output()
		bounds = image.Rect(0, 0, len(pixels[0]), len(pixels))
		img    = image.NewGray(bounds)
	)
	for posY := 0; posY < len(pixels); posY++ {
		for posX := 0; posX < len(pixels[posY]); posX++ {
			img.SetGray(posX, posY, color.Gray{Y: uint8(pixels[posY][posX])})
		}
	}
	return img
}

func ToSlice(img *image.Gray) [][]uint8 {
	var (
		M = img.Bounds().Max.X
		N = img.Bounds().Max.Y
	)

	imgSlice := ops.GeneratePixelSlice(M, N)

	for v := 0; v < N; v++ {
		for u := 0; u < M; u++ {
			imgSlice[v][u] = img.GrayAt(u, v).Y
		}
	}

	return imgSlice
}
