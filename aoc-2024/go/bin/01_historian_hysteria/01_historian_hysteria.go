package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("01_historian_hysteria")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	fst, snd, _ := parseAndSortLists(in)
	sum := 0

	for i, f := range fst {
		sum += gotils.Abs(f - snd[i])
	}

	return sum
}

func solve2(in string) int {
	fst, _, sim := parseAndSortLists(in)
	sum := 0

	for _, f := range fst {
		sum += f * sim[f]
	}

	return sum
}

func parseAndSortLists(in string) ([]int, []int, map[int]int) {
	first := make([]int, 0)
	second := make([]int, 0)
	similarity := make(map[int]int, 0)

	for i, id := range strings.Fields(in) {
		value := gotils.MustParseInt(id)

		if i%2 == 0 {
			first = append(first, value)
		} else {
			second = append(second, value)
			similarity[value] += 1
		}
	}

	slices.Sort(first)
	slices.Sort(second)

	return first, second, similarity
}
