package utils

func BorderDetection(u, v, i, j, xMax, yMax int) (int, int) {
	var (
		newU int = u + i
		newV int = v + j
	)

	if newU <= 0 {
		if newV <= 0 {
			return -newU, -newV
		} else {
			return -newU, v
		}
	}

	if newV <= 0 {
		if newU <= 0 {
			return -newU, -newV
		} else {
			return u, -newV
		}
	}

	if newU >= xMax {
		if newV >= yMax {
			return xMax - i, yMax - j
		} else {
			return xMax - i, v
		}
	}

	if newV >= yMax {
		if newU >= xMax {
			return xMax - i, yMax - j
		} else {
			return u, yMax - j
		}
	}
	return newU, newV
}
