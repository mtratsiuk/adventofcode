package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("01_secret_entrance")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0
	pos := 50
	steps := 100

	for dir, count := range rotations(in) {
		if dir == byte('L') {
			count *= -1
		}

		pos = (pos + count) % steps

		if pos < 0 {
			pos += steps
		}

		if pos == 0 {
			sum += 1
		}
	}

	return sum
}

func solve2(in string) int {
	sum := 0
	pos := 50
	steps := 100

	for dir, count := range rotations(in) {
		startPos := pos

		sum += count / steps
		count %= steps

		if dir == byte('L') {
			count *= -1
		}

		pos = pos + count

		if startPos != 0 && (pos <= 0 || pos >= steps) {
			sum += 1
		}

		pos %= steps

		if pos < 0 {
			pos += steps
		}
	}

	return sum
}

func rotations(in string) iter.Seq2[byte, int] {
	return func(yield func(byte, int) bool) {
		for line := range strings.FieldsSeq(in) {
			if !yield(line[0], gotils.MustParseInt(line[1:])) {
				return
			}
		}
	}
}
