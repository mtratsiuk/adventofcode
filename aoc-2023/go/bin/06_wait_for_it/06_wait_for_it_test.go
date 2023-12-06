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
