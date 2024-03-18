package main

import (
	"github.com/jhachmer/gocv/filter"
)

func main() {
	sharpFilter := [][]int{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	}
	/*
		gaussianFilter := [][]int{
			{1, 2, 1},
			{2, 4, 2},
			{1, 2, 1},
		}
	*/

	gauss5x5 := [][]int{
		{1, 4, 6, 4, 1},
		{4, 16, 24, 16, 4},
		{6, 24, 36, 24, 6},
		{4, 16, 24, 16, 4},
		{1, 4, 6, 4, 1},
	}

	filter.Apply2DFilter("birb.png", gauss5x5, "gauss")
	filter.Apply2DFilter("gauss.png", sharpFilter, "sharp")

}
