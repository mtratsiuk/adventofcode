package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
`

	expected := 4277556

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +
`

	expected := 3263827

	if res := solve2(strings.Trim(in, "\n")); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
