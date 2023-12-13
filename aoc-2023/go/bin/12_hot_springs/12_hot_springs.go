package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("12_hot_springs")

	fmt.Println(solve1(in))
	// fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for _, line := range strings.Split(in, "\n") {
		sum += ParseRecord(line).ArrangementsCount()
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	for _, line := range strings.Split(in, "\n") {
		sum += ParseRecordAndExpand(line).ArrangementsCount()
	}

	return sum
}

const (
	Operational = '.'
	Damaged     = '#'
	Unknown     = '?'
)

type Record struct {
	line   string
	groups []int
	known  bool
}

func ParseRecord(in string) Record {
	r := Record{}

	line, groups, _ := strings.Cut(in, " ")

	r.line = line
	r.groups = gotils.Collect(
		gotils.Map(gotils.Iter(strings.Split(groups, ",")), gotils.MustParseInt))

	return r
}

func ParseRecordAndExpand(in string) Record {
	r := Record{}

	ls, gs, _ := strings.Cut(in, " ")
	line := ls
	groups := gs

	for i := 0; i < 4; i += 1 {
		line += "?"
		line += ls

		groups += ","
		groups += gs
	}

	r.line = line
	r.groups = gotils.Collect(
		gotils.Map(gotils.Iter(strings.Split(groups, ",")), gotils.MustParseInt))

	return r
}

func (r *Record) IsKnown() bool {
	r.known = !strings.Contains(r.line, string(Unknown))
	return r.known
}

var groupToRegexp = make(map[int]*regexp.Regexp, 0)

func getRegexp(g int) *regexp.Regexp {
	r, ok := groupToRegexp[g]

	if !ok {
		r = regexp.MustCompile(fmt.Sprintf(`^[^#]*?([?#]{%v})(?:[^#]|\z)`, g))
		groupToRegexp[g] = r
	}

	return r
}

func (r Record) IsValid() bool {
	nextIdx := 0

	for _, g := range r.groups {
		if nextIdx >= len(r.line) {
			return false
		}

		singleIdx := strings.Index(r.line[nextIdx:], string(Damaged))
		idx := strings.Index(r.line[nextIdx:], strings.Repeat(string(Damaged), g))

		if idx == -1 || (r.line[singleIdx+nextIdx] == Damaged && singleIdx != idx) {
			return false
		} else {
			idx += nextIdx
		}

		noSeparatorLeft := idx != 0 && r.line[idx-1] != Operational
		noSeparatorRight := idx+g != len(r.line) && r.line[idx+g] != Operational

		if noSeparatorLeft || noSeparatorRight {
			return false
		}

		nextIdx = idx + g + 1
	}

	return nextIdx >= len(r.line) || !strings.Contains(r.line[nextIdx:], string(Damaged))
}

func (r Record) IsValidRec() bool {
	if len(r.groups) == 0 {
		return !r.known || !strings.Contains(r.line, string(Damaged))
	}

	re := getRegexp(r.groups[0])
	mIdx := re.FindStringSubmatchIndex(r.line)

	if mIdx == nil {
		return false
	}

	end := mIdx[3]

	nextRecord := Record{r.line[end:], r.groups[1:], r.known}

	if len(nextRecord.line) > 0 && nextRecord.line[0] == Unknown {
		nextRecord = nextRecord.Guess(string(Operational))
	}

	return nextRecord.IsValidRec()
}

func (r Record) Guess(option string) Record {
	return Record{strings.Replace(r.line, string(Unknown), option, 1), r.groups, r.known}
}

func (r Record) ArrangementsCount() int {
	if r.IsKnown() {
		if r.IsValidRec() {
			return 1
		}

		return 0
	}

	// lg := r.Guess(string(Operational))
	// lg.IsKnown()
	// rg := r.Guess(string(Damaged))
	// rg.IsKnown()

	// var lc, rc int

	// if lg.IsValidRec() {
	// 	lc = lg.ArrangementsCount()
	// }

	// if rg.IsValidRec() {
	// 	rc = rg.ArrangementsCount()
	// }

	// return lc + rc

	return r.Guess(string(Operational)).ArrangementsCount() + r.Guess(string(Damaged)).ArrangementsCount()
}

func (r Record) String() string {
	return fmt.Sprintf("%v %v", r.line, r.groups)
}
