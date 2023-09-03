package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/aoc-2022/go/mishutils"
)

func main() {
	in := mishutils.ReadInput("03_rucksack_reorganization")

	fmt.Println(solve_1(in))
	fmt.Println(solve_2(in))
}

func solve_1(input string) (int, error) {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		sum += getTypePriority(findWrongType(line))
	}

	return sum, nil
}

func solve_2(input string) (int, error) {
	sum := 0
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		sum += getTypePriority(findBadge(lines[i:i+3]))
	}

	return sum, nil
}

func findBadge(lines []string) rune {
	types1 := make(map[rune]struct{})
	types2 := make(map[rune]struct{})

	for _, t := range lines[0] {
		types1[t] = struct{}{}
	}

	for _, t := range lines[1] {
		if _, ok := types1[t]; ok {
			types2[t] = struct{}{}
		}
	}

	for _, t := range lines[2] {
		if _, ok := types2[t]; ok {
			return t
		}
	}

	panic("unreachable")
}

func findWrongType(line string) rune {
	types := make(map[rune]struct{})

	for _, t := range line[0 : len(line)/2] {
		types[t] = struct{}{}
	}

	for _, t := range line[len(line)/2:] {
		if _, ok := types[t]; ok {
			return t
		}
	}

	panic("unreachable")
}

const a = rune('a')
const z = rune('z')
const A = rune('A')
const Z = rune('Z')

func getTypePriority(t rune) int {
	switch {
	case t >= a && t <= z:
		return int(t-a) + 1
	case t >= A && t <= Z:
		return int(t-A) + 27
	default:
		panic("unexpected type: " + string(t))
	}
}
