package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("22_sand_slabs")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	bricks := ParseFallingBricks(in)
	slices.SortFunc(bricks, func(a, b Brick) int { return a.start.z - b.start.z })

	tower := make([]Brick, 0)

	for _, brick := range bricks {
		landingZ := 0

		for ti, tb := range tower {
			if brick.IntersectsXy(tb) {
				multipleSupports := false
				landingZ = tb.end.z

				for _, otherb := range tower[ti+1:] {
					if otherb.end.z < landingZ {
						break
					}

					if brick.IntersectsXy(otherb) {
						multipleSupports = true
					}
				}

				if !multipleSupports {
					tower[ti].disintegratable = false
				}

				break
			}
		}

		size := brick.SizeZ()
		brick.start.z = landingZ + 1
		brick.end.z = brick.start.z + size

		tower = append(tower, brick)

		slices.SortFunc(tower, func(a, b Brick) int { return b.end.z - a.end.z })
	}

	return gotils.Count(tower, func(b Brick) bool { return b.disintegratable })
}

func solve2(in string) int {
	sum := 0

	return sum
}

type Pos3d struct {
	x, y, z int
}

type Brick struct {
	start, end      Pos3d
	disintegratable bool
}

func (b Brick) String() string {
	return fmt.Sprintf("%v,%v,%v~%v,%v,%v %v",
		b.start.x, b.start.y, b.start.z,
		b.end.x, b.end.y, b.end.z,
		b.disintegratable)
}

func (b Brick) SizeZ() int {
	return b.end.z - b.start.z
}

func (b Brick) IntersectsXy(other Brick) bool {
	bxRange := Range{min(b.start.x, b.end.x), max(b.start.x, b.end.x)}
	byRange := Range{min(b.start.y, b.end.y), max(b.start.y, b.end.y)}
	oxRange := Range{min(other.start.x, other.end.x), max(other.start.x, other.end.x)}
	oyRange := Range{min(other.start.y, other.end.y), max(other.start.y, other.end.y)}
	return bxRange.OverlapsWith(oxRange) && byRange.OverlapsWith(oyRange)
}

type Range struct {
	left  int
	right int
}

func (r Range) OverlapsWith(other Range) bool {
	return other.left >= r.left && other.left <= r.right ||
		other.right <= r.right && other.right >= r.left ||
		other.Contains(r)
}

func (r Range) Contains(other Range) bool {
	return other.left >= r.left && other.right <= r.right
}

func ParseFallingBricks(in string) []Brick {
	bricks := make([]Brick, 0)

	for _, line := range strings.Split(in, "\n") {
		startStr, endStr, _ := strings.Cut(line, "~")
		start := strings.Split(startStr, ",")
		end := strings.Split(endStr, ",")

		brick := Brick{}
		brick.disintegratable = true
		brick.start.x = gotils.MustParseInt(start[0])
		brick.start.y = gotils.MustParseInt(start[1])
		brick.start.z = gotils.MustParseInt(start[2])
		brick.end.x = gotils.MustParseInt(end[0])
		brick.end.y = gotils.MustParseInt(end[1])
		brick.end.z = gotils.MustParseInt(end[2])

		if brick.start.z > brick.end.z {
			brick.start, brick.end = brick.end, brick.start
		}

		bricks = append(bricks, brick)
	}

	return bricks
}
