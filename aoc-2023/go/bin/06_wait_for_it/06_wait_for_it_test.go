package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
Time:      7  15   30
Distance:  9  40  200
`

	expected := 288

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
Time:      7  15   30
Distance:  9  40  200
`

	expected := 71503

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}

func BenchmarkWinsCount(b *testing.B) {
	race := Race{48989083, 390110311121360}

	for i := 0; i < b.N; i++ {
		race.WinsCount()
	}
}

func BenchmarkWinsCountFast(b *testing.B) {
	race := Race{48989083, 390110311121360}

	for i := 0; i < b.N; i++ {
		race.WinsCountFast()
	}
}
