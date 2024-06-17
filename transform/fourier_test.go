package transform

import (
	"testing"

	"github.com/jhachmer/imgo/img"
)

func BenchTestDataFourier() *DFT {
	inputImg := img.ReadImageFromPath("../images/Lenna.png")

	pixs := img.ToSlice(img.ConvertToGrayScale(inputImg))

	return NewDFT(pixs)
}

func BenchmarkDFT2D(b *testing.B) {
	// run the Fib function b.N times
	dft := BenchTestDataFourier()
	for n := 0; n < b.N; n++ {
		dft.DFT2D(true)
	}
}
