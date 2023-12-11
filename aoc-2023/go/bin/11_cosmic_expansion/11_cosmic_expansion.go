package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("11_cosmic_expansion")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in, 1_000_000))
}

func solve1(in string) int {
	sum := 0

	c := NewCosmos(strings.Split(in, "\n")).Expand()

	for _, pair := range gotils.Pairs(c.Galaxies()) {
		sum += gotils.Abs(pair[0].x-pair[1].x) + gotils.Abs(pair[0].y-pair[1].y)
	}

	return sum
}

func solve2(in string, expansion int) int {
	sum := 0
	c := NewCosmos(strings.Split(in, "\n"))
	colsToExpand := c.ExpandColumns()
	rowsToExpand := c.ExpandRows()

	trueX := func(x int) int {
		times := gotils.Count(colsToExpand, func(c int) bool { return c < x })
		return (expansion-1)*times + x
	}

	trueY := func(y int) int {
		times := gotils.Count(rowsToExpand, func(c int) bool { return c < y })
		return (expansion-1)*times + y
	}

	for _, pair := range gotils.Pairs(c.Galaxies()) {
		sum += gotils.Abs(trueX(pair[0].x)-trueX(pair[1].x)) + gotils.Abs(trueY(pair[0].y)-trueY(pair[1].y))
	}

	return sum
}

const (
	Galaxy = '#'
	Space  = '.'
)

type Pos2d struct {
	x, y int
}

type Cosmos struct {
	grid []string
}

func NewCosmos(grid []string) Cosmos {
	return Cosmos{grid}
}

func (c Cosmos) Galaxies() []Pos2d {
	gs := make([]Pos2d, 0)

	for y, line := range c.grid {
		for x, c := range line {
			if c == Galaxy {
				gs = append(gs, Pos2d{x, y})
			}
		}
	}

	return gs
}

func (c Cosmos) ExpandColumns() []int {
	height := len(c.grid)
	width := len(c.grid[0])
	colsToExpand := make([]int, 0)

	for x := 0; x < width; x += 1 {
		hasGalaxy := false

		for y := 0; y < height; y += 1 {
			if c.grid[y][x] == Galaxy {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			colsToExpand = append(colsToExpand, x)
		}
	}

	return colsToExpand
}

func (c Cosmos) ExpandRows() []int {
	height := len(c.grid)
	width := len(c.grid[0])
	rowsToExpand := make([]int, 0)

	for y := 0; y < height; y += 1 {
		hasGalaxy := false

		for x := 0; x < width; x += 1 {
			if c.grid[y][x] == Galaxy {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			rowsToExpand = append(rowsToExpand, y)
		}
	}

	return rowsToExpand
}

func (c Cosmos) Expand() Cosmos {
	height := len(c.grid)
	grid := make([]string, height)
	copy(grid, c.grid)
	colsToExpand := c.ExpandColumns()
	rowsToExpand := c.ExpandRows()

	for i, x := range colsToExpand {
		target := x + i

		for y := 0; y < height; y += 1 {
			grid[y] = grid[y][0:target] + string(Space) + grid[y][target:]
		}
	}

	for i, y := range rowsToExpand {
		target := y + i

		grid = append(grid[:target+1], grid[target:]...)
	}

	return Cosmos{grid}
}

func (c Cosmos) String() string {
	var sb strings.Builder

	for _, line := range c.grid {
		for _, c := range line {
			sb.WriteRune(c)
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
