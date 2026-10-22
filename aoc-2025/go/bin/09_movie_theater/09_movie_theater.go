package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("09_movie_theater")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	pairs := InputToPairs(in)

	slices.SortFunc(pairs, func(a, b Pair) int { return -cmp.Compare(a.area, b.area) })

	return pairs[0].area
}

func solve2(in string) int {
	sum := 0

	return sum
}

type Point struct {
	X, Y int
}

type Pair struct {
	left, right Point
	area        int
}

func Area(left, right Point) int {
	return gotils.Abs(left.X-right.X+1) * gotils.Abs(left.Y-right.Y+1)
}

func InputToPairs(in string) []Pair {
	points := make([]Point, 0)

	for str := range strings.FieldsSeq(in) {
		coords := strings.Split(str, ",")
		points = append(points, Point{
			gotils.MustParseInt(coords[0]),
			gotils.MustParseInt(coords[1]),
		})
	}

	pairs := make([]Pair, 0)
	pos := 0

	for i := 0; i < len(points)-1; i += 1 {
		for j := i + 1; j < len(points); j += 1 {
			left := points[i]
			right := points[j]
			pairs = append(pairs, Pair{left, right, Area(left, right)})
			pos += 1
		}
	}

	return pairs
}
