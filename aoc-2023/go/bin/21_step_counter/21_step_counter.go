package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("21_step_counter")

	fmt.Println(solve1(in, 64, Pos2d{65, 65}))
	fmt.Println(solve2(in, 256, Pos2d{65, 65}))
}

func solve1(in string, steps int, start Pos2d) int {
	grid := strings.Split(in, "\n")
	width := len(grid[0])
	height := len(grid)
	current := make(map[string]bool, 0)
	finishes := make(map[Pos2d]struct{}, 0)

	var run func(pos Pos2d, steps int)
	run = func(pos Pos2d, steps int) {
		posStepsKey := fmt.Sprintf("%v:%v:%v", pos.x, pos.y, steps)

		if pos.IsOutOfBounds(width, height) ||
			grid[pos.y][pos.x] == '#' ||
			current[posStepsKey] {
			return
		}

		if steps == 0 {
			finishes[pos] = struct{}{}
			return
		}

		current[posStepsKey] = true

		for _, m := range AllMoves90 {
			run(pos.Move(m), steps-1)
		}
	}

	run(start, steps)

	return len(finishes)
}

func solve2(in string, maxSteps int, start Pos2d) int {
	// grid := strings.Split(in, "\n")
	// width := len(grid[0])
	// height := len(grid)
	// current := make(map[string]bool, 0)
	// finishes := make(map[Pos2d]struct{}, 0)
	// finishesPerStep := make(map[int]map[Pos2d]struct{}, 0)

	// var run func(pos Pos2d, steps int)
	// run = func(pos Pos2d, steps int) {
	// 	posStepsKey := fmt.Sprintf("%v:%v:%v", pos.x, pos.y, steps)

	// 	if pos.IsOutOfBounds(width, height) ||
	// 		grid[pos.y][pos.x] == '#' ||
	// 		current[posStepsKey] {
	// 		return
	// 	}

	// 	if steps == maxSteps {
	// 		finishes[pos] = struct{}{}
	// 		return
	// 	}

	// 	if _, ok := finishesPerStep[steps]; !ok {
	// 		finishesPerStep[steps] = make(map[Pos2d]struct{}, 0)
	// 	}
	// 	finishesPerStep[steps][pos] = struct{}{}

	// 	current[posStepsKey] = true

	// 	for _, m := range AllMoves90 {
	// 		run(pos.Move(m), steps+1)
	// 	}
	// }

	// run(start, 0)

	// info := make([]string, 0)
	// for step, fps := range finishesPerStep {
	// 	info = append(info, fmt.Sprintf("[%04d]: %v", step, len(fps)))
	// }
	// slices.Sort(info)

	// fmt.Println(strings.Join(info, "\n"))
	// fmt.Println()

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
