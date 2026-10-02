package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("08_playground")

	fmt.Println(solve1(in, 1000))
	fmt.Println(solve2(in))
}

func solve1(in string, connections int) int {
	pairs := InputToPairs(in)
	slices.SortFunc(pairs, func(a, b Pair) int { return cmp.Compare(a.distance, b.distance) })

	circuits := make([]Circuit, 0)
	deleted := make([]int, 0)
	index := make(map[Point]int, 0)

	for _, pair := range pairs {
		if connections == 0 {
			break
		}

		cidxl, okl := index[pair.left]
		cidxr, okr := index[pair.right]

		if okl && !okr {
			circuits[cidxl] = append(circuits[cidxl], pair)
			index[pair.right] = cidxl
		}

		if !okl && okr {
			circuits[cidxr] = append(circuits[cidxr], pair)
			index[pair.left] = cidxr
		}

		if !okl && !okr {
			circuits = append(circuits, Circuit{pair})
			index[pair.right] = len(circuits) - 1
			index[pair.left] = len(circuits) - 1
		}

		if okl && okr && cidxl != cidxr {
			circuits[cidxl] = append(circuits[cidxl], circuits[cidxr]...)
			for _, p := range circuits[cidxr] {
				index[p.left] = cidxl
				index[p.right] = cidxl
			}
			deleted = append(deleted, cidxr)
		}

		connections -= 1
	}

	final := make([]Circuit, 0)

	for idx, c := range circuits {
		if !slices.Contains(deleted, idx) {
			final = append(final, c)
		}
	}

	slices.SortFunc(final, func(a, b Circuit) int { return cmp.Compare(len(b), len(a)) })

	return gotils.IterFluent(final[0:3]).
		Map(func(c Circuit) int { return c.Size() }).
		Fold(1, func(i1, i2 int) int { return i1 * i2 })
}

func solve2(in string) int {
	return 0
}

type Circuit []Pair

func (c Circuit) Size() int {
	s := gotils.NewSet[Point]()

	for _, p := range c {
		s.Add(p.left)
		s.Add(p.right)
	}

	return s.Size()
}

type Point struct {
	X, Y, Z int
}

type Pair struct {
	left, right Point
	distance    float64
}

func (p Point) DistanceTo(other Point) float64 {
	return math.Hypot(
		float64(p.X)-float64(other.X),
		math.Hypot(float64(p.Y)-float64(other.Y), float64(p.Z)-float64(other.Z)))
}

func (p Pair) Overlaps(other Pair) bool {
	return p.left == other.left ||
		p.left == other.right ||
		p.right == other.left ||
		p.right == other.right
}

func InputToPairs(in string) []Pair {
	points := make([]Point, 0)

	for str := range strings.FieldsSeq(in) {
		coords := strings.Split(str, ",")
		points = append(points, Point{
			gotils.MustParseInt(coords[0]),
			gotils.MustParseInt(coords[1]),
			gotils.MustParseInt(coords[2]),
		})
	}

	pairs := make([]Pair, 0)

	pos := 0

	for i := 0; i < len(points)-1; i += 1 {
		for j := i + 1; j < len(points); j += 1 {
			left := points[i]
			right := points[j]
			pairs = append(pairs, Pair{left, right, left.DistanceTo(right)})
			pos += 1
		}
	}

	return pairs
}
