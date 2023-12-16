package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("16_the_floor_will_be_lava")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	grid := strings.Split(in, "\n")
	return EnergizedFrom(grid, Pos2d{0, 0}, Move2dRight)
}

const Size = 110

func solve2(in string) int {
	grid := strings.Split(in, "\n")
	mx := math.MinInt

	for i := 0; i < Size; i += 1 {
		mx = max(
			mx,
			EnergizedFrom(grid, Pos2d{i, 0}, Move2dBottom),
			EnergizedFrom(grid, Pos2d{i, Size - 1}, Move2dTop),
			EnergizedFrom(grid, Pos2d{0, i}, Move2dRight),
			EnergizedFrom(grid, Pos2d{Size - 1, i}, Move2dLeft),
		)
	}

	return mx
}

func EnergizedFrom(grid []string, pos Pos2d, move Move2d) int {
	c := NewContraption(grid)
	c.AddBeam(pos, move)

	for c.Run() {
	}

	return len(c.energized)
}

type MirrorChar byte

const (
	MirrorCharUpwardMirror   MirrorChar = '/'
	MirrorCharDownwardMirror MirrorChar = '\\'
)

type SplitterChar byte

const (
	SplitterCharVerticalSplit   SplitterChar = '|'
	SplitterCharHorizontalSplit SplitterChar = '-'
)

const (
	CharEmpty = '.'
)

var MirrorLastMoveToNextMove = map[MirrorChar]map[Move2d]Move2d{
	MirrorCharUpwardMirror: {
		Move2dLeft:   Move2dBottom,
		Move2dRight:  Move2dTop,
		Move2dTop:    Move2dRight,
		Move2dBottom: Move2dLeft,
	},
	MirrorCharDownwardMirror: {
		Move2dLeft:   Move2dTop,
		Move2dRight:  Move2dBottom,
		Move2dTop:    Move2dLeft,
		Move2dBottom: Move2dRight,
	},
}

var SplitterLastMoveToNextMove = map[SplitterChar]map[Move2d][]Move2d{
	SplitterCharHorizontalSplit: {
		Move2dLeft:   []Move2d{Move2dLeft},
		Move2dRight:  []Move2d{Move2dRight},
		Move2dTop:    []Move2d{Move2dLeft, Move2dRight},
		Move2dBottom: []Move2d{Move2dLeft, Move2dRight},
	},
	SplitterCharVerticalSplit: {
		Move2dLeft:   []Move2d{Move2dTop, Move2dBottom},
		Move2dRight:  []Move2d{Move2dTop, Move2dBottom},
		Move2dTop:    []Move2d{Move2dTop},
		Move2dBottom: []Move2d{Move2dBottom},
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
	Move2dTop    Move2d = Move2d{0, -1}
	Move2dRight  Move2d = Move2d{1, 0}
	Move2dBottom Move2d = Move2d{0, 1}
	Move2dLeft   Move2d = Move2d{-1, 0}
)

type Beam struct {
	pos      Pos2d
	lastMove Move2d
}

func (b *Beam) Move(move Move2d) {
	b.lastMove = move
	b.pos.x += move.dx
	b.pos.y += move.dy
}

type Contraption struct {
	grid      []string
	energized map[Pos2d][]Move2d
	beams     []*Beam
}

func NewContraption(grid []string) Contraption {
	c := Contraption{}
	c.grid = grid
	c.energized = make(map[Pos2d][]Move2d, 0)
	c.beams = make([]*Beam, 0)

	return c
}

func (c *Contraption) AddBeam(pos Pos2d, move Move2d) *Beam {
	b := &Beam{pos, move}
	c.beams = append(c.beams, b)
	return b
}

func (c *Contraption) Run() bool {
	running := false

	for i := range c.beams {
		b := c.beams[i]

		if ms, ok := c.energized[b.pos]; ok && slices.Contains(ms, b.lastMove) {
			continue
		}

		moved := c.MoveBeam(b)

		if moved {
			running = true
		}
	}

	return running
}

func (c *Contraption) MoveBeam(beam *Beam) bool {
	if beam.pos.x < 0 ||
		beam.pos.x >= c.Width() ||
		beam.pos.y < 0 ||
		beam.pos.y >= c.Height() {
		return false
	}

	c.energized[beam.pos] = append(c.energized[beam.pos], beam.lastMove)

	cur := c.grid[beam.pos.y][beam.pos.x]

	switch {
	case cur == CharEmpty:
		beam.Move(beam.lastMove)
	case cur == byte(MirrorCharDownwardMirror) || cur == byte(MirrorCharUpwardMirror):
		beam.Move(MirrorLastMoveToNextMove[MirrorChar(cur)][beam.lastMove])
	case cur == byte(SplitterCharHorizontalSplit) || cur == byte(SplitterCharVerticalSplit):
		moves := SplitterLastMoveToNextMove[SplitterChar(cur)][beam.lastMove]

		if len(moves) > 1 {
			newBeam := c.AddBeam(beam.pos, beam.lastMove)
			newBeam.Move(moves[1])
		}

		beam.Move(moves[0])
	default:
		panic("unexpected char")
	}

	return true
}

func (c *Contraption) Width() int {
	return len(c.grid[0])
}

func (c *Contraption) Height() int {
	return len(c.grid)
}
