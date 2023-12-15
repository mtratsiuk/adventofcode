package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("15_lens_library")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for _, step := range strings.FieldsFunc(in, func(r rune) bool { return r == ',' }) {
		sum += hash(step)
	}

	return sum
}

func solve2(in string) int {
	sum := 0
	boxes := make([][]Lens, 256)

	for _, step := range strings.FieldsFunc(in, func(r rune) bool { return r == ',' }) {
		lens := ParseLens(step)
		boxIdx := hash(lens.label)
		lensIdx := slices.IndexFunc(boxes[boxIdx], func(l Lens) bool { return l.label == lens.label })

		if lens.operation == OpDash {
			if lensIdx == -1 {
				continue
			}

			boxes[boxIdx] = append(boxes[boxIdx][:lensIdx], boxes[boxIdx][lensIdx+1:]...)
		}

		if lens.operation == OpEq {
			if lensIdx == -1 {
				boxes[boxIdx] = append(boxes[boxIdx], lens)
				continue
			}

			boxes[boxIdx][lensIdx].length = lens.length
		}
	}

	for bi, box := range boxes {
		for li, lens := range box {
			sum += (1 + bi) * (1 + li) * lens.length
		}
	}

	return sum
}

func hash(in string) int {
	val := 0

	for _, c := range in {
		val += int(c)
		val *= 17
		val %= 256
	}

	return val
}

const (
	OpDash = string('-')
	OpEq   = string('=')
)

type Lens struct {
	label     string
	operation string
	length    int
}

func ParseLens(in string) Lens {
	var label, operation, length string

	if strings.Contains(in, OpEq) {
		label, length, _ = strings.Cut(in, OpEq)
		operation = OpEq
	} else {
		label, _ = strings.CutSuffix(in, OpDash)
		operation = OpDash
		length = "0"
	}

	return Lens{label, operation, gotils.MustParseInt(length)}
}
