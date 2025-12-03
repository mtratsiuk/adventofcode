package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

	expected := 3

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

	expected := 6

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
