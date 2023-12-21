package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`

	expected := 16

	if res := solve1(strings.TrimSpace(in), 6, Pos2d{5, 5}); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`

	expected := 0

	if res := solve2(strings.TrimSpace(in), 64, Pos2d{5, 5}); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
