package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
3   4
4   3
2   5
1   3
3   9
3   3
`

	expected := 11

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
3   4
4   3
2   5
1   3
3   9
3   3
`

	expected := 31

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
