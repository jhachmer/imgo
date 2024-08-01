package ascii

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jhachmer/imgo/internal/ops"
)

func ToAscii(in [][]uint8) [][]string {
	rows, cols := len(in), len(in[0])
	ret := ops.GenerateSlice[string](cols, rows)
	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			c := in[y][x]
			level := c / 51
			if level == 5 {
				level--
			}
			ret[y][x] = levels[level]
		}
	}
	return ret
}

func WriteAscii(in [][]string, outPath string) {
	rows, cols := len(in), len(in[0])
	f, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for y := 0; y < rows; y++ {
		var s strings.Builder
		for x := 0; x < cols; x++ {
			_, err := fmt.Fprintf(&s, "%v", in[y][x])
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err := fmt.Fprintln(f, s.String())
		if err != nil {
			log.Fatal(err)
		}
	}
}
