package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("13_point_of_incidence")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for _, pattern := range strings.Split(in, "\n\n") {
		p := NewPattern(strings.Split(pattern, "\n"))

		for i := 0; i < p.Width()-1; i += 1 {
			if ReflectedAt(p.columns, i, 0) {
				sum += i + 1
			}
		}

		for i := 0; i < p.Height()-1; i += 1 {
			if ReflectedAt(p.rows, i, 0) {
				sum += 100 * (i + 1)
			}
		}
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	return sum
}

func ReflectedAt(lines []string, n, offset int) bool {
	left := n - offset
	right := n + 1 + offset

	if left < 0 || right >= len(lines) {
		return true
	}

	if lines[left] != lines[right] {
		return false
	}

	return ReflectedAt(lines, n, offset+1)
}

type Pattern struct {
	grid    []string
	rows    []string
	columns []string
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

func (p *Pattern) Row(n int) string {
	return p.grid[n]
}

func (p *Pattern) Column(n int) string {
	return p.columns[n]
}

func (p *Pattern) Width() int {
	return len(p.grid[0])
}

func (p *Pattern) Height() int {
	return len(p.grid)
}
