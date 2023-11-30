package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("04_camp_cleanup")

	fmt.Println(solve_1(in))
	fmt.Println(solve_2(in))
}

func solve_1(input string) (int, error) {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		a, b := RangePairFromString(line)

		if a.Contains(b) || b.Contains(a) {
			count += 1
		}
	}

	return count, nil
}

func solve_2(input string) (int, error) {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		a, b := RangePairFromString(line)

		if a.OverlapsWith(b) {
			count += 1
		}
	}

	return count, nil
}

type Range struct {
	left  int
	right int
}

func (r Range) Contains(other Range) bool {
	return other.left >= r.left && other.right <= r.right
}

func (r Range) OverlapsWith(other Range) bool {
	return other.left >= r.left && other.left <= r.right ||
		other.right <= r.right && other.right >= r.left ||
		other.Contains(r)
}

func RangeFromString(s string) Range {
	rangeStr := strings.Split(s, "-")

	l := gotils.MustParseInt(rangeStr[0])
	r := gotils.MustParseInt(rangeStr[1])

	return Range{
		l,
		r,
	}
}

func RangePairFromString(s string) (Range, Range) {
	pair := strings.Split(s, ",")

	first := RangeFromString(pair[0])
	second := RangeFromString(pair[1])

	return first, second
}
