package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

	expected := 114

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

	expected := 2

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
