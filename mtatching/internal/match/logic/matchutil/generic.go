package matchutil

type number interface {
	int | int32 | int64 | float32 | float64
}

func Abs[T number](num T) T {
	if num > 0 {
		return num
	}
	return -num
}

func Min[T number](a, b T) T {
	if a > b {
		return b
	}
	return a
}
