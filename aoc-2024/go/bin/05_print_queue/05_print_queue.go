package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("05_print_queue")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	rules, updates := parseRules(in)
	sum := 0

	for _, update := range updates {
		if rules.IsCorrect(update) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func solve2(in string) int {
	rules, updates := parseRules(in)
	sum := 0

	for _, update := range updates {
		if !rules.IsCorrect(update) {
			slices.SortFunc(update, func(a, b int) int {
				if gotils.MapContains(rules[a].before, b) || gotils.MapContains(rules[b].after, a) {
					return 1
				}
				return -1
			})

			sum += update[len(update)/2]
		}
	}

	return sum
}

type Rule struct {
	before, after map[int]struct{}
}

func NewRule() Rule {
	return Rule{make(map[int]struct{}), make(map[int]struct{})}
}

type Rules map[int]Rule
type Update []int

func (r Rules) IsCorrect(u Update) bool {
	for i, page := range u {
		correctBefore := gotils.Every(u[:i], func(v int) bool { return gotils.MapContains(r[page].before, v) })
		correctAfter := gotils.Every(u[i+1:], func(v int) bool { return gotils.MapContains(r[page].after, v) })

		if !correctBefore || !correctAfter {
			return false
		}
	}

	return true
}

func parseRules(in string) (Rules, []Update) {
	rules := make(Rules, 0)
	updates := make([]Update, 0)

	rulesStr, updatesStr, _ := strings.Cut(in, "\n\n")

	for _, line := range strings.Split(rulesStr, "\n") {
		fstStr, sndStr, _ := strings.Cut(line, "|")
		fst, snd := gotils.MustParseInt(fstStr), gotils.MustParseInt(sndStr)

		if _, ok := rules[fst]; !ok {
			rules[fst] = NewRule()
		}

		if _, ok := rules[snd]; !ok {
			rules[snd] = NewRule()
		}

		rules[fst].after[snd] = struct{}{}
		rules[snd].before[fst] = struct{}{}
	}

	for _, line := range strings.Split(updatesStr, "\n") {
		updates = append(
			updates,
			gotils.Collect(gotils.Map(gotils.Iter(strings.Split(line, ",")), gotils.MustParseInt)),
		)
	}

	return rules, updates
}
