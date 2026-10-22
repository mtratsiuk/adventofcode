package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("07_bridge_repair")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for l := range strings.SplitSeq(in, "\n") {
		line := ParseEqLine(l)

		if line.CanBeMadeTrue(false) {
			sum += line.value
		}
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	for l := range strings.SplitSeq(in, "\n") {
		line := ParseEqLine(l)

		if line.CanBeMadeTrue(true) {
			sum += line.value
		}
	}

	return sum
}

type EqLine struct {
	value   int
	numbers []int
}

func ParseEqLine(line string) EqLine {
	val, numsStr, _ := strings.Cut(line, ": ")

	return EqLine{
		value:   gotils.MustParseInt(val),
		numbers: gotils.IterFluentSeq(strings.SplitSeq(numsStr, " ")).Map(gotils.MustParseInt).Collect(),
	}
}

func (line *EqLine) CanBeMadeTrue(withConcat bool) bool {
	variants := []int{line.numbers[0]}

	for _, num := range line.numbers[1:] {
		next := make([]int, 0)

		for _, variant := range variants {
			next = append(next, variant+num)
			next = append(next, variant*num)

			if withConcat {
				next = append(next, gotils.MustParseInt(fmt.Sprintf("%d%d", variant, num)))
			}
		}

		variants = next
	}

	return slices.Contains(variants, line.value)
}
