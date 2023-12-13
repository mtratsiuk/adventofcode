package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	in := `
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

	expected := 21

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

	expected := 21

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}

func TestIsValidRec(t *testing.T) {
	in := `
?###???????? 3,2,1
`

	expected := true

	if res := ParseRecord(strings.TrimSpace(in)).IsValidRec(); res != expected {
		t.Errorf("TestIsValidRec() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
