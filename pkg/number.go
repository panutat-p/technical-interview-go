package pkg

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func INT(str string) int {
	r, _ := strconv.Atoi(str)
	return r
}

func INT64(str string) int64 {
	r, _ := strconv.ParseInt(str, 10, 64)
	return r
}
