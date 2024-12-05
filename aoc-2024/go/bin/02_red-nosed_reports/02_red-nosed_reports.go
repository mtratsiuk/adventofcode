package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("02_red-nosed_reports")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	reports := parseReports(in)
	sum := 0

	for _, report := range reports {
		if isSafe(report) {
			sum += 1
		}
	}

	return sum
}

func solve2(in string) int {
	reports := parseReports(in)
	sum := 0

	for _, report := range reports {
		if isSafe(report) {
			sum += 1
		} else {
			for i := range report {
				if isSafe(append(slices.Clone(report[:i]), report[i+1:]...)) {
					sum += 1
					break
				}
			}
		}
	}

	return sum
}

func isSafe(report []int) bool {
	mind := 1
	maxd := 3

	wasInc := report[0] < report[1]

	for i := 1; i < len(report); i += 1 {
		f, s := report[i-1], report[i]
		dt := gotils.Abs(f - s)
		isInc := f < s

		if wasInc != isInc || dt < mind || dt > maxd {
			return false
		}

		wasInc = isInc
	}

	return true
}

func parseReports(in string) [][]int {
	reports := make([][]int, 0)

	for _, line := range strings.Split(in, "\n") {
		reports = append(
			reports,
			gotils.Collect(gotils.Map(gotils.Iter(strings.Fields(line)), gotils.MustParseInt)),
		)
	}

	return reports
}
