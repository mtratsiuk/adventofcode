package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("05_if_you_give_a_seed_a_fertilizer")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	locs := make([]int, 0)
	mappings := make([]Mappings, 0)
	seedsStr, mappingsStr, _ := strings.Cut(in, "\n\n")
	_, seeds, _ := strings.Cut(seedsStr, ": ")

	for _, m := range strings.Split(mappingsStr, "\n\n") {
		mappings = append(mappings, ParseMappings(m))
	}

	for _, seed := range strings.Fields(seeds) {
		next := gotils.MustParseInt(seed)

		for _, ms := range mappings {
			next = ms.Map(next)
		}

		locs = append(locs, next)
	}

	return gotils.Min(locs)
}

func solve2(in string) int {
	sum := 0

	return sum
}

type Mapping struct {
	src, dst, rng int
}

type Mappings []Mapping

func (mpgs Mappings) Map(val int) int {
	for _, m := range mpgs {
		if val >= m.src && val <= m.src+m.rng {
			step := val - m.src
			return m.dst + step
		}
	}

	return val
}

func ParseMappings(lines string) []Mapping {
	mappings := make([]Mapping, 0)
	mStrings := strings.Split(lines, "\n")[1:]

	for _, mStr := range mStrings {
		mappingStr := strings.Fields(mStr)
		m := Mapping{}
		m.dst = gotils.MustParseInt(mappingStr[0])
		m.src = gotils.MustParseInt(mappingStr[1])
		m.rng = gotils.MustParseInt(mappingStr[2])
		mappings = append(mappings, m)
	}

	return mappings
}
