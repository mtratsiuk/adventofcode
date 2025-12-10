package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("06_trash_compactor")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	operators, operands := parse(in)
	results := make([]int, len(operators))

	for _, line := range operands {
		for i, operand := range line {
			next := gotils.MustParseInt(operand)
			op := operators[i]

			if op == "+" {
				results[i] += next
			} else {
				if results[i] == 0 {
					results[i] = next
				} else {
					results[i] *= next
				}
			}
		}
	}

	return gotils.Sum(results)
}

func solve2(in string) int {
	sum := 0

	lines := strings.Split(in, "\n")
	operators := lines[len(lines)-1]
	operands := lines[:len(lines)-1]

	for len(operators) < len(operands[0]) {
		operators += " "
	}

	start := len(operators) - 1
	end := start

	for start >= 0 {
		op := operators[start]

		if operators[start] == ' ' {
			start -= 1
			continue
		}

		result := 0

		for end >= start {
			next := ""

			for _, line := range operands {
				next += string(line[end])
			}

			nextInt := gotils.MustParseInt(strings.TrimSpace(next))

			if op == '+' {
				result += nextInt
			} else {
				if result == 0 {
					result = nextInt
				} else {
					result *= nextInt
				}
			}

			end -= 1
		}

		sum += result
		start -= 2
		end = start
	}

	return sum
}

func parse(in string) ([]string, [][]string) {
	lines := strings.Split(in, "\n")
	operators := strings.Fields(lines[len(lines)-1])
	operands := gotils.Collect(gotils.Map(gotils.Iter(lines[0:len(lines)-1]), strings.Fields))

	return operators, operands
}
