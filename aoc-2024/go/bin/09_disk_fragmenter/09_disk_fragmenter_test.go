package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
2333133121414131402
`

	expected := 1928

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
2333133121414131402
`

	expected := 2858

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
