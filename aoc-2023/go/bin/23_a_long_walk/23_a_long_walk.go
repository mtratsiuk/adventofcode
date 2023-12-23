package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("23_a_long_walk")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	grid := strings.Split(in, "\n")
	width := len(grid[0])
	height := len(grid)
	start := Pos2d{1, 0}
	end := Pos2d{width - 2, height - 1}

	maxSteps := 0

	var run func(pos, lastPos Pos2d, steps int)
	run = func(pos, lastPos Pos2d, steps int) {
		if pos.IsOutOfBounds(width, height) ||
			grid[pos.y][pos.x] == '#' {
			return
		}

		if pos == end {
			maxSteps = max(maxSteps, steps)
			return
		}

		for _, m := range AllMoves90 {
			nextPos := pos.Move(m)

			if nextPos == lastPos ||
				(grid[pos.y][pos.x] != '.' && m != ArrowToMove2d[grid[pos.y][pos.x]]) {
				continue
			}

			run(nextPos, pos, steps+1)
		}
	}

	run(start, start, 0)

	return maxSteps
}

func solve2(in string) int {
	return 0
}

type Pos2d struct {
	x, y int
}

func (p Pos2d) IsOutOfBounds(width, height int) bool {
	return p.x < 0 || p.x >= width || p.y < 0 || p.y >= height
}

func (p Pos2d) Move(m Move2d) Pos2d {
	return Pos2d{p.x + m.dx, p.y + m.dy}
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

var AllMoves90 = []Move2d{
	Move2dTop,
	Move2dRight,
	Move2dBottom,
	Move2dLeft,
}

var ArrowToMove2d = map[byte]Move2d{
	'>': Move2dRight,
	'v': Move2dBottom,
	'<': Move2dLeft,
	'^': Move2dTop,
}
