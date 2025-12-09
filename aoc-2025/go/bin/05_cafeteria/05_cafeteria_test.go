package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

	expected := 3

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

	expected := 14

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
