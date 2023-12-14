package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

	expected := 136

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

	expected := 64

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}

func TestTiltLine(t *testing.T) {
	cases := []struct{ in, expected string }{
		{`OO.O.O..##`, `OOOO....##`},
		{`OO.O.O..`, `OOOO....`},
		{`#`, `#`},
		{`OO..##..O#`, `OO..##O..#`},
		{`..OO`, `OO..`},
	}

	for _, c := range cases {
		if res := TiltLine(strings.TrimSpace(c.in)); res != c.expected {
			t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", c.expected, res)
		}
	}
}

func TestEqual(t *testing.T) {
	in := `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

	p1 := NewPattern(strings.Split(strings.TrimSpace(in), "\n"))
	p2 := p1.TiltCycle()

	if res := p1.Equals(&p2); res != false {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", false, res)
	}

	if res := p1.Equals(&p1); res != true {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", true, res)
	}
}
