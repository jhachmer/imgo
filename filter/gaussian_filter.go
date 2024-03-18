package filter

import (
	"fmt"
	"math"
)

func CreateGaussKernel1D(sigma float64) []float64 {
	center := int(3.0 * sigma)
	fmt.Printf("center %d\n", center)
	h := make([]float64, 2*center+1)
	sigma2 := sigma * sigma
	for i := range h {
		r := float64(center - i)
		h[i] = math.Exp(-0.5 * (r * r) / sigma2)
	}
	return h
}
