package utils

func ClampPixel(value, upper, lower int) int {
	if value < lower {
		value = lower
	}
	if value > upper {
		value = upper
	}
	return value
}

func GeneratePixelSlice(x, y int) [][]uint8 {
	res := make([][]uint8, y)
	for i := 0; i < y; i++ {
		res[i] = make([]uint8, x)
	}
	return res
}
