package main

import (
	"testing"
)

func Test1(t *testing.T) {
	in := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	expected := 142

	if res := solve1(in); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	expected := 281

	if res := solve2(in); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
