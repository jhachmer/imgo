package mathlib

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type NumberUnsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type NumberSigned interface {
	~int | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T NumberSigned](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Max[T NumberSigned | NumberUnsigned](a []T) T {
	var m T
	for i := range a {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func Min[T NumberSigned | NumberUnsigned](a []T) T {
	var m = a[0]
	for i := range a {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}

func Sum[T NumberSigned](values []T, abs bool) T {
	var sum T
	for _, v := range values {
		if abs {
			sum += Abs(v)
		} else {
			sum += v
		}
	}
	return sum
}
