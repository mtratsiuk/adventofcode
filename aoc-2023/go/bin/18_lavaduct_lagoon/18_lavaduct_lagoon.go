package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("18_lavaduct_lagoon")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	lines := strings.Split(in, "\n")
	path := make(map[Pos2d]*PathNode, 0)

	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	pos := Pos2d{0, 0}

	var first, prev *PathNode

	for _, l := range lines {
		ds := ParseDigStep(l)
		move := DirToMove[ds.dir]

		for i := 0; i < ds.len; i += 1 {
			pos.x += move.dx
			pos.y += move.dy

			minX = min(minX, pos.x)
			minY = min(minY, pos.y)
			maxX = max(maxX, pos.x)
			maxY = max(maxY, pos.y)

			node := PathNode{pos, ds, i == ds.len-1, prev, nil}

			if prev != nil {
				prev.next = &node
			} else {
				first = &node
			}

			path[pos] = &node
			prev = &node
		}
	}

	prev.next = first
	first.prev = prev

	ray := func(x, y int) int {
		intersections := 0
		lastIntersectionDir := 0

		for i := minX; i <= x; i += 1 {
			pos := Pos2d{i, y}
			node, found := path[pos]

			if found && (!IsHorizontalDir(node.ds.dir) || node.edge) {
				intrDir := 0

				if node.prev.pos.y < pos.y || node.next.pos.y > pos.y {
					intrDir = 1
				}

				if node.prev.pos.y > pos.y || node.next.pos.y < pos.y {
					intrDir = -1
				}

				if lastIntersectionDir != intrDir {
					intersections += 1
					lastIntersectionDir = intrDir
				}
			}
		}

		return intersections
	}

	edge := len(path)
	interior := 0

	for y := minY; y <= maxY; y += 1 {
		for x := minX; x <= maxX; x += 1 {
			if _, found := path[Pos2d{x, y}]; ray(x, y)%2 == 1 && !found {
				interior += 1
			}
		}
	}

	return edge + interior
}

func solve2(in string) int {
	lines := strings.Split(in, "\n")
	vertsHorizontal := make([]Vert, 0)
	vertsVertical := make([]Vert, 0)

	perimeter := 0
	start := Pos2d{0, 0}

	for _, l := range lines {
		ds := ParseDigStepHex(l)
		perimeter += ds.len
		move := DirToMove[ds.dir]
		end := Pos2d{start.x + move.dx*ds.len, start.y + move.dy*ds.len}

		if IsHorizontalDir(ds.dir) {
			vertsHorizontal = append(vertsHorizontal, Vert{start, end})
		} else {
			vertsVertical = append(vertsVertical, Vert{start, end})
		}

		start = end
	}

	slices.SortFunc(vertsHorizontal, func(a, b Vert) int { return a.start.y - b.start.y })
	slices.SortFunc(vertsVertical, func(a, b Vert) int { return a.start.x - b.start.x })

	minY := vertsHorizontal[0].start.y
	maxY := gotils.Last(vertsHorizontal).start.y

	square := 0

	for y := minY; y <= maxY; y += 1 {
		intersections := 0
		lastIntersectionDir := 0

		prevVerticalIdx := -1

		for idx, vertical := range vertsVertical {
			if !vertical.ContainsY(y) {
				continue
			}

			if prevVerticalIdx != -1 {
				prevVertical := vertsVertical[prevVerticalIdx]
				testVert := Vert{Pos2d{prevVertical.start.x, 0}, Pos2d{vertical.start.x, 0}}

				onHorizontalEdge := slices.IndexFunc(vertsHorizontal, func(v Vert) bool {
					return v.start.y == y && (testVert.ContainsX(v.start.x) && testVert.ContainsX(v.end.x))
				}) != -1

				if intersections%2 == 1 && !onHorizontalEdge {
					square += gotils.Abs(vertical.start.x-prevVertical.start.x) - 1
				}
			}

			intersectionDir := 0

			if vertical.start.y > vertical.end.y {
				intersectionDir = -1
			} else {
				intersectionDir = 1
			}

			if lastIntersectionDir != intersectionDir {
				intersections += 1
				lastIntersectionDir = intersectionDir
			}

			prevVerticalIdx = idx
		}
	}

	return square + perimeter
}

func (v *Vert) ContainsX(x int) bool {
	mx := max(v.start.x, v.end.x)
	mn := min(v.start.x, v.end.x)
	return x >= mn && x <= mx
}

func (v *Vert) ContainsY(y int) bool {
	mx := max(v.start.y, v.end.y)
	mn := min(v.start.y, v.end.y)
	return y >= mn && y <= mx
}

type Vert struct {
	start, end Pos2d
}

var DirToMove = map[byte]Move2d{
	'U': Move2dTop,
	'D': Move2dBottom,
	'L': Move2dLeft,
	'R': Move2dRight,
}

func IsHorizontalDir(dir byte) bool {
	return dir == 'L' || dir == 'R'
}

type Pos2d struct {
	x, y int
}

type Move2d struct {
	dx, dy int
}

var (
	Move2dTop    Move2d = Move2d{0, -1}
	Move2dRight  Move2d = Move2d{1, 0}
	Move2dBottom Move2d = Move2d{0, 1}
	Move2dLeft   Move2d = Move2d{-1, 0}
)

type PathNode struct {
	pos  Pos2d
	ds   DigStep
	edge bool
	prev *PathNode
	next *PathNode
}

type DigStep struct {
	dir byte
	len int
	hex string
}

func ParseDigStep(line string) DigStep {
	ds := DigStep{}

	chunks := strings.Split(line, " ")

	ds.dir = chunks[0][0]
	ds.len = gotils.MustParseInt(chunks[1])
	ds.hex = chunks[2][2 : len(chunks[2])-1]

	return ds
}

var CodeToDir = map[byte]byte{
	'0': 'R',
	'1': 'D',
	'2': 'L',
	'3': 'U',
}

func ParseDigStepHex(line string) DigStep {
	ds := ParseDigStep(line)
	ds.dir = CodeToDir[ds.hex[len(ds.hex)-1:][0]]

	if len, err := strconv.ParseInt(ds.hex[:len(ds.hex)-1], 16, 64); err == nil {
		ds.len = int(len)
	} else {
		panic(err)
	}

	return ds
}
