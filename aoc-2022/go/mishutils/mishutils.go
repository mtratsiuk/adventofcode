package mishutils

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

func Max[T cmp.Ordered](s []T) (max T, ok bool) {
	if len(s) == 0 {
		ok = false
		return
	}

	ok = true
	max = s[0]

	for _, cur := range s[1:] {
		if cur > max {
			max = cur
		}
	}

	return
}

func Sum[T constraints.Integer | constraints.Float](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}

	return
}
