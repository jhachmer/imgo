package transform

import (
	"github.com/jhachmer/imgo/img"
	"testing"
)

func BenchTestDataHough() *HoughTransform {
	inputImg := img.ReadImageFromPath("../images/Lenna.png")

	pixs := img.ToSlice(img.ConvertToGrayScale(inputImg))

	return newHoughTransform(pixs, 500, 500)
}

func BenchmarkHoughLines(b *testing.B) {
	// run the Fib function b.N times
	hough := BenchTestDataHough()
	for n := 0; n < b.N; n++ {
		HoughLines(hough.Input, 500, 500)
	}
}
