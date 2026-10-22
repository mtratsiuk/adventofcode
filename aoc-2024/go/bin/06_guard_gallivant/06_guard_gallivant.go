package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("06_guard_gallivant")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	visited := make(map[gotils.Pos2d]struct{}, 0)
	grid := strings.Split(in, "\n")
	flatGrid := strings.Join(grid, "")
	height := len(grid)
	width := len(grid[0])

	posIdx := strings.IndexFunc(flatGrid, func(r rune) bool { return CharToDirection(r) != nil })
	gotils.Assert(posIdx != -1, func() string { return fmt.Sprintf("failed to find start pos: %v", posIdx) })

	pos := gotils.NewPos2d(posIdx%width, posIdx/width)
	dir := *CharToDirection(rune(flatGrid[posIdx]))

	for !pos.IsOutOfBounds(width, height) {
		visited[pos] = struct{}{}
		next := pos.Move(dir)

		if next.IsOutOfBounds(width, height) {
			break
		}

		for grid[next.Y][next.X] == '#' {
			dir = dir.NextDirection90Clockwise()
			next = pos.Move(dir)
		}

		pos = next
	}

	return len(visited)
}

func solve2(in string) int {
	sum := 0

	return sum
}

func CharToDirection(c rune) *gotils.Pos2d {
	switch c {
	case '^':
		return &gotils.DirN
	case '>':
		return &gotils.DirE
	case 'v':
		return &gotils.DirS
	case '<':
		return &gotils.DirW
	default:
		return nil
	}
}
