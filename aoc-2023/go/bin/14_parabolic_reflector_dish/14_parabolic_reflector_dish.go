package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("14_parabolic_reflector_dish")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	p := NewPattern(strings.Split(in, "\n"))

	for _, column := range p.columns {
		sum += Score(TiltLine(column))
	}

	return sum
}

const CycleTimes = 1_000_000_000

func solve2(in string) int {
	p := NewPattern(strings.Split(in, "\n"))
	cycles := make([]Pattern, 0)

	for i := 0; i < CycleTimes; i += 1 {
		cycles = append(cycles, p)
		p = p.TiltCycle()

		for j, c := range cycles {
			if c.Equals(&p) {
				loop := cycles[j:]
				cyclesLeft := CycleTimes - j
				target := cyclesLeft % len(loop)
				lastP := loop[target]

				sum := 0

				for _, column := range lastP.columns {
					sum += Score(column)
				}

				return sum
			}
		}
	}

	panic("unreachable")
}

type Pattern struct {
	grid    []string
	rows    []string
	columns []string
}

const (
	Empty = '.'
	Round = 'O'
	Cube  = '#'
)

func Score(line string) int {
	score := 0

	for i, c := range line {
		if c == Round {
			score += len(line) - i
		}
	}

	return score
}

func TiltLine(line string) string {
	cubeIdx := strings.IndexRune(line, Cube)
	cubeChar := string(Cube)
	next := ""

	if cubeIdx == -1 {
		cubeIdx = len(line)
		cubeChar = ""
	}

	if cubeIdx < len(line)-1 {
		next = TiltLine(line[cubeIdx+1:])
	}

	roundCount := strings.Count(line[:cubeIdx], string(Round))
	emptyCount := cubeIdx - roundCount

	return strings.Repeat(string(Round), roundCount) +
		strings.Repeat(string(Empty), emptyCount) +
		cubeChar +
		next
}

func (p *Pattern) TiltCycle() Pattern {
	tiltedNorth := PatternFromColumns(gotils.Collect(gotils.Map(gotils.Iter(p.columns),
		func(x string) string { return TiltLine(x) })))

	tiltedWest := NewPattern(gotils.Collect(gotils.Map(gotils.Iter(tiltedNorth.rows),
		func(x string) string { return TiltLine(x) })))

	tiltedSouth := PatternFromColumns(gotils.Collect(gotils.Map(gotils.Iter(tiltedWest.columns),
		func(x string) string { return gotils.ReverseAscii(TiltLine(gotils.ReverseAscii(x))) })))

	tiltedEast := NewPattern(gotils.Collect(gotils.Map(gotils.Iter(tiltedSouth.rows),
		func(x string) string { return gotils.ReverseAscii(TiltLine(gotils.ReverseAscii(x))) })))

	return tiltedEast
}

func NewPattern(grid []string) Pattern {
	p := Pattern{}
	p.grid = grid
	p.rows = make([]string, p.Height())
	p.columns = make([]string, p.Width())

	for y := 0; y < p.Height(); y += 1 {
		p.rows[y] = grid[y]

		for x := 0; x < p.Width(); x += 1 {
			p.columns[x] += string(grid[y][x])
		}
	}

	return p
}

func PatternFromColumns(columns []string) Pattern {
	width := len(columns)
	height := len(columns[0])

	grid := make([]string, height)

	for y := 0; y < height; y += 1 {
		var row strings.Builder

		for x := 0; x < width; x += 1 {
			row.WriteByte(columns[x][y])
		}

		grid[y] = row.String()
	}

	return NewPattern(grid)
}

func (p *Pattern) Width() int {
	return len(p.grid[0])
}

func (p *Pattern) Height() int {
	return len(p.grid)
}

func (p *Pattern) Equals(other *Pattern) bool {
	for y, line := range p.grid {
		if line != other.grid[y] {
			return false
		}
	}

	return true
}

func (p Pattern) String() string {
	var sb strings.Builder

	for _, line := range p.grid {
		for _, c := range line {
			sb.WriteRune(c)
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
