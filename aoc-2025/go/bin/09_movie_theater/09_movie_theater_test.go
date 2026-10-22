package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

	expected := 50

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
data
`

	expected := 0

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
