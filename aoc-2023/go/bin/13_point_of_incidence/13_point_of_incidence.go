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
		sum += GetReflectionScores(&p)[0]
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	for _, pattern := range strings.Split(in, "\n\n") {
		p := NewPattern(strings.Split(pattern, "\n"))
		sum += GetSmudgedScore(&p)
	}

	return sum
}

func GetSmudgedScore(p *Pattern) int {
	oldScores := GetReflectionScores(p)

	if len(oldScores) > 1 {
		panic("should not happen")
	}

	for y := 0; y < p.Height(); y += 1 {
		for x := 0; x < p.Width(); x += 1 {
			np := p.WithSmudgeAt(x, y)
			newScores := GetReflectionScores(&np)

			for _, score := range newScores {
				if score != oldScores[0] {
					return score
				}
			}
		}
	}

	panic("unreachable")
}

func GetReflectionScores(p *Pattern) []int {
	scores := make([]int, 0)

	for i := 0; i < p.Width()-1; i += 1 {
		if ReflectedAt(p.columns, i, 0) {
			scores = append(scores, i+1)
		}
	}

	for i := 0; i < p.Height()-1; i += 1 {
		if ReflectedAt(p.rows, i, 0) {
			scores = append(scores, 100*(i+1))
		}
	}

	return scores
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

func (p *Pattern) WithSmudgeAt(x, y int) Pattern {
	np := Pattern{}
	np.grid = p.grid
	np.rows = make([]string, p.Height())
	copy(np.rows, p.rows)
	np.columns = make([]string, p.Width())
	copy(np.columns, p.columns)

	var smudge rune

	if np.grid[y][x] == '.' {
		smudge = '#'
	} else {
		smudge = '.'
	}

	np.rows[y] = np.rows[y][:x] + string(smudge) + np.rows[y][x+1:]
	np.grid[y] = np.rows[y]
	np.columns[x] = np.columns[x][:y] + string(smudge) + np.columns[x][y+1:]

	return np
}

func (p *Pattern) Width() int {
	return len(p.grid[0])
}

func (p *Pattern) Height() int {
	return len(p.grid)
}
