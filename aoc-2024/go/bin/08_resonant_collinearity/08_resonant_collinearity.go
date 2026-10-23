package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("08_resonant_collinearity")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	uniqueAntinodes := make(map[gotils.Pos2d]struct{}, 0)
	grid := NewGrid(in)

	for _, points := range grid.antennas {
		for _, pair := range gotils.Pairs(points) {
			pairVec := gotils.NewVec2(pair[0], pair[1])

			for _, vec := range []gotils.Vec2{pairVec, pairVec.Reverse()} {
				antinodePos := vec.Scale(2).To
				if !antinodePos.IsOutOfBounds(grid.width, grid.height) {
					uniqueAntinodes[antinodePos] = struct{}{}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func solve2(in string) int {
	uniqueAntinodes := make(map[gotils.Pos2d]struct{}, 0)
	grid := NewGrid(in)

	for _, points := range grid.antennas {
		for _, pair := range gotils.Pairs(points) {
			pairVec := gotils.NewVec2(pair[0], pair[1])

			for _, vec := range []gotils.Vec2{pairVec, pairVec.Reverse()} {
				norm := vec.NormaliseInt()
				dx, dy := norm.Sides()

				for !norm.To.IsOutOfBounds(grid.width, grid.height) {
					uniqueAntinodes[norm.To] = struct{}{}
					norm = norm.Add(dx, dy)
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

type Grid struct {
	grid          [][]rune
	width, height int
	antennas      map[rune][]gotils.Pos2d
}

func NewGrid(in string) Grid {
	grid := make([][]rune, 0)
	antennas := make(map[rune][]gotils.Pos2d, 0)

	for y, line := range strings.Fields(in) {
		lineSlice := make([]rune, 0)

		for x, char := range strings.Split(line, "") {
			c := rune(char[0])
			lineSlice = append(lineSlice, c)

			if c == '.' {
				continue
			}

			antennas[c] = append(antennas[c], gotils.NewPos2d(x, y))
		}

		grid = append(grid, lineSlice)
	}

	return Grid{
		grid:     grid,
		width:    len(grid[0]),
		height:   len(grid),
		antennas: antennas,
	}
}
