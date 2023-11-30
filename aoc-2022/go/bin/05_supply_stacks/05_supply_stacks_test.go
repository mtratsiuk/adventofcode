package main

import (
	"testing"
)

func Test1(t *testing.T) {
	in := `
data
`

	expected := 0

	if res := solve1(in); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
data
`

	expected := 0

	if res := solve2(in); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
