package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("17_clumsy_crucible")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	return solve17(in, 3, 0)
}

func solve2(in string) int {
	return solve17(in, 10, 4)
}

func solve17(in string, MaxStepsWithoutTurn, MinStepsWithoutTurn int) int {
	grid := ParseIntGrid(in)
	width := len(grid[0])
	height := len(grid)
	heatLossMap := gotils.Make2d[int](width, height, math.MaxInt)

	var run func(x, y, steps, loss, attempts int, lastMove Move2d, fx, fy int)
	run = func(x, y, steps, loss, attempts int, lastMove Move2d, fx, fy int) {
		// exit: went out of bounds
		if x < 0 || x >= width || y < 0 || y >= height {
			return
		}

		// exit: should've turned
		if steps > MaxStepsWithoutTurn {
			return
		}

		// exit: can't stop in the end
		if steps < MinStepsWithoutTurn && x == width-1 && y == height-1 {
			return
		}

		curLoss := loss + grid[y][x]

		// exit: there was shorter path here
		if curLoss >= heatLossMap[y][x] {
			if attempts >= MaxStepsWithoutTurn+MinStepsWithoutTurn {
				return
			} else {
				attempts += 1
			}
		} else {
			attempts = 0
		}

		// exit: got too far back
		maxLookBehind := max(MaxStepsWithoutTurn+MinStepsWithoutTurn, 10)
		if gotils.Abs(fx-x)+gotils.Abs(fy-y) > maxLookBehind {
			return
		}

		heatLossMap[y][x] = min(heatLossMap[y][x], curLoss)
		fx = max(fx, x)
		fy = max(fy, y)

		for _, move := range AllMoves90 {
			stepsCounter := steps + 1
			nx, ny := x+move.dx, y+move.dy

			// can't go back
			if nx+lastMove.dx == x && ny+lastMove.dy == y {
				continue
			}

			// can't turn yet
			if steps < MinStepsWithoutTurn && move != lastMove {
				continue
			}

			if move != lastMove {
				stepsCounter = 1
			}

			run(nx, ny, stepsCounter, curLoss, attempts, move, fx, fy)
		}
	}

	for _, move := range AllMoves90 {
		run(move.dx, move.dy, 1, 0, 0, move, 0, 0)
	}

	return heatLossMap[height-1][width-1]
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

func ParseIntGrid(in string) [][]int {
	lines := strings.Split(in, "\n")
	height := len(lines)
	grid := make([][]int, height)

	for y, line := range lines {
		grid[y] = gotils.Collect(gotils.Map(gotils.Iter(strings.Split(line, "")), gotils.MustParseInt))
	}

	return grid
}
