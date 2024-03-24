package main

import (
	"github.com/jhachmer/gocv/filter"
	"github.com/jhachmer/gocv/utils"
)

func main() {

	//sharpFilter := [][]int{
	//	{0, -1, 0},
	//	{-1, 5, -1},
	//	{0, -1, 0},
	//}
	///*
	//	gaussianFilter := [][]int{
	//		{1, 2, 1},
	//		{2, 4, 2},
	//		{1, 2, 1},
	//	}
	//*/
	//
	//gauss5x5 := [][]int{
	//	{1, 4, 6, 4, 1},
	//	{4, 16, 24, 16, 4},
	//	{6, 24, 36, 24, 6},
	//	{4, 16, 24, 16, 4},
	//	{1, 4, 6, 4, 1},
	//}

	inputImg := utils.ConvertImageToGrayPNG("Lenna.png")
	gradientSlice := filter.SobelOperator(inputImg)

	imgGradX, imgGradY := utils.CreateNewGrayFromGradient(gradientSlice)

	utils.WriteGrayToFilePNG("xgrad", imgGradX)
	utils.WriteGrayToFilePNG("ygrad", imgGradY)

}
