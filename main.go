package main

import (
	"github.com/jhachmer/gocv/edge"
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
	gauss5x5 := [][]int{
		{1, 4, 7, 4, 1},
		{4, 16, 26, 16, 4},
		{7, 26, 41, 26, 7},
		{4, 16, 26, 16, 4},
		{1, 4, 7, 4, 1},
	}

	inputImg := utils.ConvertImageToGrayPNG("Camera_obscura.png")
	blurredImg := filter.Apply2DFilter(inputImg, gauss5x5)

	utils.WriteGrayToFilePNG("input", inputImg)
	utils.WriteGrayToFilePNG("gauss", blurredImg)

	gradientSlice := edge.SobelOperator(blurredImg)

	magPixels := edge.CalcMagnitudeFromGradient(gradientSlice)
	grayMag := utils.CreateNewMagnitudeFromGradient(magPixels)

	utils.WriteGrayToFilePNG("magpix", grayMag)

	imgGradX, imgGradY := utils.CreateNewGrayFromGradient(gradientSlice)
	utils.WriteGrayToFilePNG("xgrad", imgGradX)
	utils.WriteGrayToFilePNG("ygrad", imgGradY)

}
