package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
987654321111111
811111111111119
234234234234278
818181911112111
`

	expected := 357

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
987654321111111
811111111111119
234234234234278
818181911112111
`

	expected := 3121910778619

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
