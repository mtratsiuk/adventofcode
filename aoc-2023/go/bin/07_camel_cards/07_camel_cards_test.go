package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

	expected := 6440

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

	expected := 5905

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
