package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/mtratsiuk/adventofcode/aoc-2022/go/mishutils"
)

func main() {
    fmt.Println(solve_1(readInput("01_calorie_counting")))
    fmt.Println(solve_2(readInput("01_calorie_counting")))
}

func solve_1(input string) (int, error) {
    elves, err := calories(input)

    if err != nil {
        return 0, err
    }

    max, ok := mishutils.Max(elves)

    if !ok {
        return max, errors.New("failed to find max calories")
    }

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

func readInput(name string) string {
    pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	input, err := os.ReadFile(path.Join(pwd, "data", name + ".txt"))

    if err != nil {
		panic(err)
	}

    return string(input)
}
