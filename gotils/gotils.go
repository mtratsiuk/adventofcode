package gotils

import (
	"fmt"
	"math"
	"slices"
	"strconv"

	"golang.org/x/exp/constraints"
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

func Abs[T Numeric](v T) T {
	if v < 0 {
		return -v
	}

	return v
}

func Sum[T Numeric](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}

	return
}

func Every[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

func Count[T any](s []T, f func(T) bool) int {
	count := 0

	for _, v := range s {
		if f(v) {
			count += 1
		}
	}

	return count
}

func Gcd[T constraints.Integer](a, b T) T {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func GcdAll[T constraints.Integer](arr []T) T {
	gcd := arr[0]

	for _, n := range arr[1:] {
		gcd = Gcd(gcd, n)
	}

	return gcd
}

func Lcm[T constraints.Integer](a, b T) T {
	return (a * b) / Gcd(a, b)
}

func LcmAll[T constraints.Integer](arr []T) T {
	lcm := arr[0]

	for _, n := range arr[1:] {
		lcm = Lcm(lcm, n)
	}

	return lcm
}

func Combine[T any](slices [][]T) [][]T {
	next := make([][]T, 0)

	for _, x := range slices[0] {
		next = append(next, []T{x})
	}

	for _, x := range slices[1:] {
		next = combine(next, x)
	}

	return next
}

func combine[T any](left [][]T, right []T) [][]T {
	res := make([][]T, 0)

	for _, l := range left {
		res = append(res, combineInner(l, right)...)
	}

	return res
}

func combineInner[T any](left, right []T) [][]T {
	res := make([][]T, 0)

	for _, r := range right {
		re := make([]T, 0)
		re = append(re, left...)
		re = append(re, r)
		res = append(res, re)
	}

	return res
}

func Pairs[T any](values []T) [][]T {
	pairs := make([][]T, 0)

	for i, l := range values {
		for _, r := range values[i+1:] {
			pairs = append(pairs, []T{l, r})
		}
	}

	return pairs
}

func ReverseAscii(in string) string {
	r := []rune(in)
	slices.Reverse(r)
	return string(r)
}

func Make2d[T any](width, height int, init T) [][]T {
	grid := make([][]T, height)

	for y := 0; y < height; y += 1 {
		grid[y] = make([]T, width)

		for x := 0; x < width; x += 1 {
			grid[y][x] = init
		}
	}

	return grid
}

func Last[T any](s []T) T {
	return s[len(s)-1]
}
