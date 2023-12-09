package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("09_mirage_maintenance")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for _, line := range strings.Split(in, "\n") {
		seq := gotils.Collect(gotils.Map(gotils.Iter(strings.Fields(line)), gotils.MustParseInt))
		last := make([]int, 0)

		for {
			last = append(last, seq[len(seq)-1])

			if gotils.Every(seq, func(i int) bool { return i == 0 }) {
				break
			}

			seq = nextSeq(seq)
		}

		sum += gotils.Sum(last)
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	for _, line := range strings.Split(in, "\n") {
		seq := gotils.Collect(gotils.Map(gotils.Iter(strings.Fields(line)), gotils.MustParseInt))
		first := make([]int, 0)

		for {
			first = append(first, seq[0])

			if gotils.Every(seq, func(i int) bool { return i == 0 }) {
				break
			}

			seq = nextSeq(seq)
		}

		slices.Reverse(first)
		sum += gotils.Fold(gotils.Iter(first), 0, func(i1, i2 int) int { return i2 - i1 })
	}

	return sum
}

func nextSeq(seq []int) []int {
	next := make([]int, 0)

	for i := range seq[1:] {
		next = append(next, seq[i+1]-seq[i])
	}

	return next
}
