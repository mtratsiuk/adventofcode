package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("06_wait_for_it")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	res := 1
	races := ParseRaces(in)

	for _, r := range races {
		res *= r.WinsCount()
	}

	return res
}

func solve2(in string) int {
	return ParseRace(in).WinsCount()
}

type Race struct {
	time, dist int
}

func (r Race) WinsCount() int {
	count := 0

	for ht := 0; ht <= r.time; ht += 1 {
		v := ht
		t := r.time - ht
		d := v * t

		if d > r.dist {
			count += 1
		}
	}

	return count
}

func ParseRaces(in string) []Race {
	races := make([]Race, 0)

	timesL, distL, _ := strings.Cut(in, "\n")
	times := strings.Fields(strings.TrimLeftFunc(timesL, func(r rune) bool { return !unicode.IsDigit(r) }))
	dist := strings.Fields(strings.TrimLeftFunc(distL, func(r rune) bool { return !unicode.IsDigit(r) }))

	for i, t := range times {
		races = append(races, Race{gotils.MustParseInt(t), gotils.MustParseInt(dist[i])})
	}

	return races
}

func ParseRace(in string) Race {
	re := regexp.MustCompile(`\D`)

	timesL, distL, _ := strings.Cut(in, "\n")
	timeS := re.ReplaceAllString(timesL, "")
	distS := re.ReplaceAllString(distL, "")

	return Race{gotils.MustParseInt(timeS), gotils.MustParseInt(distS)}
}
