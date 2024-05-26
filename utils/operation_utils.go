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

