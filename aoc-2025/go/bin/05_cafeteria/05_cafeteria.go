package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("05_cafeteria")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	ranges, ids := parse(in)

	for id := range ids {
		idi := gotils.MustParseInt(id)

		if gotils.Any(ranges, func(r *gotils.Pos2d) bool { return idi >= r.X && idi <= r.Y }) {
			sum += 1
		}
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	ranges, _ := parse(in)

	for _, r := range ranges {
		sum += r.Y - r.X + 1
	}

	return sum
}

func parse(in string) ([]*gotils.Pos2d, iter.Seq[string]) {
	rangesStr, idsStr, _ := strings.Cut(in, "\n\n")
	ranges := make([]*gotils.Pos2d, 0)

	for r := range strings.FieldsSeq(rangesStr) {
		l, r, _ := strings.Cut(r, "-")
		rng := gotils.NewPos2d(gotils.MustParseInt(l), gotils.MustParseInt(r))
		ranges = append(ranges, &rng)
	}

	slices.SortFunc(ranges, func(a, b *gotils.Pos2d) int { return a.X - b.X })

	for i := 0; i < len(ranges)-1; {
		if ranges[i+1].X <= ranges[i].Y {
			ranges[i].Y = max(ranges[i].Y, ranges[i+1].Y)
			ranges = slices.Delete(ranges, i+1, min(i+2, len(ranges)))
		} else {
			i += 1
		}
	}

	return ranges, strings.FieldsSeq(idsStr)
}
