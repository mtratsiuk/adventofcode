package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mtratsiuk/adventofcode/aoc-2022/go/mishutils"
)

func main() {
	in := mishutils.ReadInput("04_camp_cleanup")

	fmt.Println(solve_1(in))
	fmt.Println(solve_2(in))
}

func solve_1(input string) (int, error) {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		a, b, err := RangePairFromString(line)

		if err != nil {
			return 0, err
		}

		if a.Contains(b) || b.Contains(a) {
			count += 1
		}
	}

	return count, nil
}

func solve_2(input string) (int, error) {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		a, b, err := RangePairFromString(line)

		if err != nil {
			return 0, err
		}

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

func RangeFromString(s string) (Range, error) {
	rangeStr := strings.Split(s, "-")

	l, err := strconv.Atoi(rangeStr[0])

	if err != nil {
		return Range{}, err
	}

	r, err := strconv.Atoi(rangeStr[1])

	if err != nil {
		return Range{}, err
	}

	return Range{
		l,
		r,
	}, nil
}

func RangePairFromString(s string) (Range, Range, error) {
	pair := strings.Split(s, ",")

	first, err := RangeFromString(pair[0])

	if err != nil {
		return Range{}, Range{}, err
	}

	second, err := RangeFromString(pair[1])

	if err != nil {
		return Range{}, Range{}, err
	}

	return first, second, nil
}
