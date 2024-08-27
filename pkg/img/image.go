package img

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/jhachmer/imgo/internal/ops"
)

type ImageType interface {
	ImageGray
}

type Outputer interface {
	Output() [][]uint8
}

type OutputFunc func() [][]uint8

type ImageGray struct {
	Pixels [][]uint8
}

func NewImageGray(path string) (*ImageGray, error) {
	pix, err := FileToSliceGray(path)
	if err != nil {
		return nil, err
	}
	return &ImageGray{
		Pixels: pix,
	}, nil
}

func (i *ImageGray) Output() [][]uint8 {
	return i.Pixels
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
func ReadImageFromPath(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sourceImg, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return sourceImg, nil
}

func FileToSliceGray(path string) ([][]uint8, error) {
	img, err := ReadImageFromPath(path)
	if err != nil {
		return nil, err
	}
	gray := ConvertToGrayScale(img)
	return ToSlice(gray), nil
}

// ToPNG  writes Go image to filesystem
// Creates image at given filepath
func ToPNG(outputFileName string, newImage *image.Gray) error {
	f, err := os.Create(outputFileName + ".png")
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, newImage); err != nil {
		return err
	}
	return nil
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
			img.SetGray(posX, posY, color.Gray{Y: pixels[posY][posX]})
		}
	}
	return img
}

func ToSlice(img *image.Gray) [][]uint8 {
	var (
		M = img.Bounds().Max.X
		N = img.Bounds().Max.Y
	)

	imgSlice := ops.GenerateSlice[uint8](M, N)

	for v := 0; v < N; v++ {
		for u := 0; u < M; u++ {
			imgSlice[v][u] = img.GrayAt(u, v).Y
		}
	}

	return imgSlice
}
