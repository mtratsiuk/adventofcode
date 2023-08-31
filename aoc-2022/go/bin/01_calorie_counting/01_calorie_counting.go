package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mtratsiuk/adventofcode/aoc-2022/go/mishutils"
)

func main() {
	in := mishutils.ReadInput("01_calorie_counting")

	fmt.Println(solve_1(in))
	fmt.Println(solve_2(in))
}

func solve_1(input string) (int, error) {
	elves, err := calories(input)

	if err != nil {
		return 0, err
	}

	max := mishutils.Max(elves)

	return max, nil
}

func solve_2(input string) (int, error) {
	elves, err := calories(input)

	if err != nil {
		return 0, err
	}

	sort.IntSlice(elves).Sort()

	return mishutils.Sum(elves[len(elves)-3:]), nil
}

func calories(input string) ([]int, error) {
	elves := make([]int, 10)
	cur := 0

	for _, line := range strings.Split(input, "\n") {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			elves = append(elves, cur)
			cur = 0
		} else {
			c, err := strconv.Atoi(trimmed)

			if err != nil {
				return nil, err
			}

			cur += c
		}
	}

	return elves, nil
}
