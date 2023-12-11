package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

var StartPos = Pos2d{18, 112}

func main() {
	in := gotils.ReadInput("10_pipe_maze")

	fmt.Println(solve1(in, StartPos))
	fmt.Println(solve2(in, StartPos))
}

func solve1(in string, start Pos2d) int {
	loopSteps := make(map[Pos2d]int, 0)
	maze := NewMaze(strings.Split(in, "\n"))

	maze.pos = start
	maze.Move(Move2dRight)

	for maze.Current() != DirCharStart {
		loopSteps[maze.pos] = maze.steps
		maze.Step()
	}

	maze.pos = start
	maze.steps = 0
	maze.Move(Move2dBottom)

	for {
		if maze.steps == loopSteps[maze.pos] {
			return maze.steps
		}

		maze.Step()
	}
}

func solve2(in string, start Pos2d) int {
	count := 0
	grid := strings.Split(in, "\n")
	width := len(grid[0])
	height := len(grid)
	pipe := make(map[Pos2d]bool, 0)
	pipePrevs := make(map[Pos2d]Pos2d, 0)
	pipeNexts := make(map[Pos2d]Pos2d, 0)

	maze := NewMaze(grid)
	prev := start
	pipe[start] = true
	maze.pos = start
	maze.Move(Move2dBottom)

	for maze.Current() != DirCharStart {
		pipe[maze.pos] = true
		pipePrevs[maze.pos] = prev
		pipeNexts[prev] = maze.pos
		prev = maze.pos
		maze.Step()
	}

	ray := func(x, y int) int {
		intscts := 0
		lastIntrDir := 0

		for i := 0; i <= x; i += 1 {
			pos := Pos2d{i, y}

			if pipe[pos] && grid[pos.y][pos.x] != byte(DirCharEastWest) {
				intrDir := 0

				if pipePrevs[pos].y < pos.y || pipeNexts[pos].y > pos.y {
					intrDir = 1
				}

				if pipePrevs[pos].y > pos.y || pipeNexts[pos].y < pos.y {
					intrDir = -1
				}

				if lastIntrDir != intrDir {
					intscts += 1
					lastIntrDir = intrDir
				}
			}

			if !pipe[pos] {
				if intscts%2 == 0 {
					lastIntrDir = 0
				}
			}
		}

		return intscts
	}

	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			if ray(x, y)%2 == 1 && !pipe[Pos2d{x, y}] {
				count += 1
			}
		}
	}

	return count
}

type DirChar byte

const (
	DirCharNorthSouth DirChar = '|'
	DirCharEastWest   DirChar = '-'
	DirCharNorthEast  DirChar = 'L'
	DirCharNorthWest  DirChar = 'J'
	DirCharSouthWest  DirChar = '7'
	DirCharSouthEast  DirChar = 'F'
	DirCharGround     DirChar = '.'
	DirCharStart      DirChar = 'S'
)

var DirCharLastMoveToNextMove = map[DirChar]map[Move2d]Move2d{
	DirCharNorthSouth: {
		Move2dBottom: Move2dBottom,
		Move2dTop:    Move2dTop,
	},
	DirCharEastWest: {
		Move2dRight: Move2dRight,
		Move2dLeft:  Move2dLeft,
	},
	DirCharNorthEast: {
		Move2dBottom: Move2dRight,
		Move2dLeft:   Move2dTop,
	},
	DirCharNorthWest: {
		Move2dBottom: Move2dLeft,
		Move2dRight:  Move2dTop,
	},
	DirCharSouthWest: {
		Move2dTop:   Move2dLeft,
		Move2dRight: Move2dBottom,
	},
	DirCharSouthEast: {
		Move2dTop:  Move2dRight,
		Move2dLeft: Move2dBottom,
	},
}

type Pos2d struct {
	x, y int
}

type Move2d struct {
	dx int
	dy int
}

var (
	Move2dTop         Move2d = Move2d{0, -1}
	Move2dRight       Move2d = Move2d{1, 0}
	Move2dTopRight    Move2d = Move2d{1, -1}
	Move2dTopLeft     Move2d = Move2d{-1, -1}
	Move2dBottom      Move2d = Move2d{0, 1}
	Move2dLeft        Move2d = Move2d{-1, 0}
	Move2dBottomRight Move2d = Move2d{1, 1}
	Move2dBottomLeft  Move2d = Move2d{-1, 1}
)

var AllMoves = []Move2d{
	Move2dTop,
	Move2dRight,
	Move2dTopRight,
	Move2dTopLeft,
	Move2dBottom,
	Move2dLeft,
	Move2dBottomRight,
	Move2dBottomLeft,
}

var AllMoves90 = []Move2d{
	Move2dTop,
	Move2dRight,
	Move2dBottom,
	Move2dLeft,
}

type Maze struct {
	grid     []string
	pos      Pos2d
	lastMove Move2d
	steps    int
}

func NewMaze(grid []string) Maze {
	m := Maze{}
	m.grid = grid
	return m
}

func (m *Maze) Current() DirChar {
	return DirChar(m.grid[m.pos.y][m.pos.x])
}

func (m *Maze) Move(move Move2d) {
	m.lastMove = move
	m.pos.x = m.pos.x + move.dx
	m.pos.y = m.pos.y + move.dy
	m.steps += 1
}

func (m *Maze) Step() {
	move := DirCharLastMoveToNextMove[m.Current()][m.lastMove]
	m.Move(move)
}
