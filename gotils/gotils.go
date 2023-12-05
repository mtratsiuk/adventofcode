package gotils

import (
	"fmt"
	"math"
	"strconv"
)

func MustParseInt(s string) int {
	r, err := strconv.Atoi(s)

	if err != nil {
		panic(fmt.Sprintf("ParseInt: %v", err))
	}

	return r
}

func Max[T Numeric](s []T) T {
	if len(s) == 0 {
		return T(math.NaN())
	}

	max := s[0]

	for _, cur := range s[1:] {
		if cur > max {
			max = cur
		}
	}

	return max
}

func Min[T Numeric](s []T) T {
	if len(s) == 0 {
		return T(math.NaN())
	}

	min := s[0]

	for _, cur := range s[1:] {
		if cur < min {
			min = cur
		}
	}

	return min
}

func Sum[T Numeric](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}

	return
}
