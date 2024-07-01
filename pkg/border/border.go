package border

// Detection returns tuple of updated pixel position based on the image bounds
// (u,v) is initial position
// (iOffset, jOffset) are the current position in applied filter kernel
// (xMax, yMax) are the maximum pixel length in each direction
func Detection(u, v, iOffset, jOffset, xMax, yMax int) (int, int) {
	newU := u + iOffset
	newV := v + jOffset

	// Mirror newU
	if newU < 0 {
		newU = -newU
	} else if newU > xMax {
		newU = 2*xMax - newU + 1
	}

	// Mirror newV
	if newV < 0 {
		newV = -newV
	} else if newV > yMax {
		newV = 2*yMax - newV + 1
	}

	return newU, newV
}
