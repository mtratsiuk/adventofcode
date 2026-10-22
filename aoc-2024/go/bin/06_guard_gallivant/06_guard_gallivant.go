package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

const OBSTACLE_RUNE = '#'

func main() {
	in := gotils.ReadInput("06_guard_gallivant")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	grid := NewGrid(in)
	visited, _ := grid.Walk()

	return len(visited)
}

func solve2(in string) int {
	sum := 0
	grid := NewGrid(in)
	startPos, startDir := grid.pos, grid.dir

	visited, _ := grid.Walk()

	for pathPos := range visited {
		if pathPos == startPos {
			continue
		}

		grid.pos = startPos
		grid.dir = startDir

		grid.grid[pathPos.Y][pathPos.X] = OBSTACLE_RUNE

		if _, ok := grid.Walk(); !ok {
			sum += 1
		}

		grid.grid[pathPos.Y][pathPos.X] = '.'
	}

	return sum
}

type Grid struct {
	grid          [][]rune
	width, height int
	pos, dir      gotils.Pos2d
}

func NewGrid(in string) Grid {
	grid := strings.Split(in, "\n")
	flatGrid := strings.Join(grid, "")
	height := len(grid)
	width := len(grid[0])

	posIdx := strings.IndexFunc(flatGrid, func(r rune) bool { return CharToDirection(r) != nil })
	gotils.Assert(posIdx != -1, func() string { return fmt.Sprintf("failed to find start pos: %v", posIdx) })

	pos := gotils.NewPos2d(posIdx%width, posIdx/width)
	dir := *CharToDirection(rune(flatGrid[posIdx]))

	g := Grid{
		grid:   gotils.IterFluent(grid).Map(func(s string) []rune { return []rune(s) }).Collect(),
		width:  width,
		height: height,
		pos:    pos,
		dir:    dir,
	}

	return g
}

func (g *Grid) Walk() (map[gotils.Pos2d]gotils.Pos2d, bool) {
	visited := make(map[gotils.Pos2d]gotils.Pos2d, 0)

	for !g.pos.IsOutOfBounds(g.width, g.height) {
		visited[g.pos] = g.dir
		next := g.pos.Move(g.dir)

		if next.IsOutOfBounds(g.width, g.height) {
			break
		}

		for g.grid[next.Y][next.X] == OBSTACLE_RUNE {
			g.dir = g.dir.NextDirection90Clockwise()
			next = g.pos.Move(g.dir)
		}

		if dir, ok := visited[next]; ok && dir == g.dir {
			return visited, false
		}

		g.pos = next
	}

	return visited, true
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
