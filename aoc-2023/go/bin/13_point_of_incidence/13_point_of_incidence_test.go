package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

	expected := 405

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

	expected := 400

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}

func TestWithSmudgeAt(t *testing.T) {
	in := `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.
`

	expectedRow := `###.....#`
	expectedCol := `###.###`

	p := NewPattern(strings.Split(strings.TrimSpace(in), "\n"))
	np := p.WithSmudgeAt(2, 2)

	if res := np.rows[2]; res != expectedRow {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expectedRow, res)
	}

	if res := np.columns[2]; res != expectedCol {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expectedCol, res)
	}
}
