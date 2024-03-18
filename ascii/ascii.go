package ascii

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func PrintGray(path string) {
	catFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	img, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(img)

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}
