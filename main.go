package main

import (
	"fmt"

	"github.com/jhachmer/gocv/edge"
	"github.com/jhachmer/gocv/filter"
	"github.com/jhachmer/gocv/model"
	"github.com/jhachmer/gocv/utils"
)

func main() {

	outFolder := "out/"

	gauss7x7 := [][]int{
		{0, 0, 1, 2, 1, 0, 0},
		{0, 3, 13, 22, 13, 3, 0},
		{1, 13, 59, 97, 59, 13, 1},
		{2, 22, 97, 159, 97, 22, 2},
		{1, 13, 59, 97, 59, 13, 1},
		{0, 3, 13, 22, 13, 3, 0},
		{0, 0, 1, 2, 1, 0, 0},
	}

	k7x7 := model.NewKernel2D(gauss7x7)

	inputImg := utils.ConvertImageToGrayPNG("images/siwi.png")
	blurredImg := filter.Apply2DFilter(inputImg, *k7x7)

	utils.WriteGrayToFilePNG(outFolder+"input", inputImg)
	utils.WriteGrayToFilePNG(outFolder+"gauss", blurredImg)

	gradientSlice, err := edge.SobelOperator(blurredImg)
	if err != nil {
		panic(err)
	}

	magPixels := edge.BuildGradientMagnitudeSlice(gradientSlice, edge.SOBEL_COEFF_SUM)

	canny := edge.CannyEdgeDetector(gradientSlice, magPixels, 15, 60)
	cannyImage := utils.SliceToImage(canny)
	utils.WriteGrayToFilePNG("out/canny", cannyImage)
	grayMag := utils.CreateMagnitudeGrayImageFromGradient(magPixels)

	utils.WriteGrayToFilePNG(outFolder+"magpix", grayMag)

	//imgGradX, imgGradY := utils.CreateXAndYGradientGrayImage(gradientSlice)
	//utils.WriteGrayToFilePNG(outFolder+"xgrad", imgGradX)
	//utils.WriteGrayToFilePNG(outFolder+"ygrad", imgGradY)

	fmt.Println("Irgendwas")
}
