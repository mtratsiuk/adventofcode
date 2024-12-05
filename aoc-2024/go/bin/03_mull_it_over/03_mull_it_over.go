package main

import (
	"fmt"
	"regexp"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("03_mull_it_over")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	mulPattern := regexp.MustCompile(`mul\((?<first>\d+),(?<second>\d+)\)`)
	sum := 0

	for _, match := range mulPattern.FindAllStringSubmatch(in, -1) {
		sum += gotils.MustParseInt(match[1]) * gotils.MustParseInt(match[2])
	}

	return sum
}

func solve2(in string) int {
	mulPattern := regexp.MustCompile(`(mul\((?<first>\d+),(?<second>\d+)\))|(do\(\))|(don't\(\))`)
	sum := 0
	do := true

	for _, match := range mulPattern.FindAllStringSubmatch(in, -1) {
		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		} else if do {
			sum += gotils.MustParseInt(match[2]) * gotils.MustParseInt(match[3])
		}
	}

	return sum
}
