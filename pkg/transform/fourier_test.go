package transform

import (
	"testing"

	"github.com/jhachmer/imgo/pkg/img"
)

func BenchTestDataFourier() *DFT {
	inputImg, err := img.ReadImageFromPath("../images/Lenna.png")
	if err != nil {
		panic(err)
	}

	pixs := img.ToSlice(img.ConvertToGrayScale(inputImg))

	return NewDFT(pixs)
}

func BenchmarkDFT2D(b *testing.B) {
	// run the Fib function b.N times
	dft := BenchTestDataFourier()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		DFT2D(dft.Transformed, true)
	}
}
