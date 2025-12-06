package main

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("03_lobby")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for bank := range batteries(in) {
		sum += maxJoltage(bank, 2)
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	for bank := range batteries(in) {
		sum += maxJoltage(bank, 12)
	}

	return sum
}

func maxJoltage(bank []int, digits int) int {
	joltage := 0

	for digits > 0 {
		a := slices.Max(bank[0 : len(bank)-(digits-1)])
		joltage += a * int(math.Pow10(digits-1))
		bank = bank[slices.Index(bank, a)+1:]
		digits -= 1
	}

	return joltage
}

func batteries(in string) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for line := range strings.FieldsSeq(in) {
			if !yield(gotils.Collect(gotils.Map(gotils.Iter(strings.Split(line, "")), gotils.MustParseInt))) {
				return
			}
		}
	}
}
