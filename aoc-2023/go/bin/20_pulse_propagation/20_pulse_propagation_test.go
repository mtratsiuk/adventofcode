package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`

	expected := 32000000

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
