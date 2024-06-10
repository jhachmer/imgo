package util

func BorderDetection(u, v, iOffset, jOffset, xMax, yMax int) (int, int) {
	var (
		newU = u + iOffset
		newV = v + jOffset
	)

	if newU <= 0 {
		if newV <= 0 {
			return -newU, -newV
		} else {
			return -newU, newV
		}
	}

	if newV <= 0 {
		return newU, -newV
	}

	if newU >= xMax {
		if newV >= yMax {
			return xMax - iOffset, yMax - jOffset
		} else {
			return xMax - iOffset, newV
		}
	}

	if newV >= yMax {
		return newU, yMax - jOffset
	}
	return newU, newV
}
