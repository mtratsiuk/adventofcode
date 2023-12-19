package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("19_aplenty")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	wsStr, partsStr, _ := strings.Cut(in, "\n\n")
	ws := ParseWorkflows(strings.Split(wsStr, "\n"))
	parts := strings.Split(partsStr, "\n")

	results := make(chan int, len(parts))
	sum := 0

	for _, p := range parts {
		go WalkWorkflows(ws, ParsePart(p), results)
	}

	for i := 0; i < len(parts); i += 1 {
		sum += <-results
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	return sum
}

func WalkWorkflows(ws Workflows, part Part, c chan int) {
	rules := ws["in"]

	for {
		for _, r := range rules {
			if !r.test(part) {
				continue
			}

			if r.dest == "R" {
				c <- 0
				return
			}

			if r.dest == "A" {
				c <- part.Value()
				return
			}

			rules = ws[r.dest]
			break
		}
	}
}

func less(key string, threshold int) func(Part) bool {
	return func(p Part) bool {
		return p[key] < threshold
	}
}

func more(key string, threshold int) func(Part) bool {
	return func(p Part) bool {
		return p[key] > threshold
	}
}

func always(p Part) bool {
	return true
}

type Rule struct {
	test func(Part) bool
	dest string
}

type Part map[string]int

func (p Part) Value() int {
	sum := 0

	for _, c := range p {
		sum += c
	}

	return sum
}

type Workflows map[string][]Rule

func ParseWorkflows(lines []string) Workflows {
	ws := make(Workflows, 0)

	for _, line := range lines {
		rules := make([]Rule, 0)
		name, rulesStr, _ := strings.Cut(line, "{")

		for _, ruleStr := range strings.Split(rulesStr[:len(rulesStr)-1], ",") {
			test, dest, ok := strings.Cut(ruleStr, ":")
			var rule Rule

			if !ok {
				rule.test = always
				rule.dest = test
			} else {
				key := string(test[0])
				op := test[1]
				threshold := gotils.MustParseInt(test[2:])

				rule.dest = dest

				if op == '<' {
					rule.test = less(key, threshold)
				} else {
					rule.test = more(key, threshold)
				}
			}

			rules = append(rules, rule)
		}

		ws[name] = rules
	}

	return ws
}

var partRe = regexp.MustCompile(`\d+`)

func ParsePart(line string) Part {
	part := make(Part, 0)

	m := partRe.FindAllString(line, -1)

	part["x"] = gotils.MustParseInt(m[0])
	part["m"] = gotils.MustParseInt(m[1])
	part["a"] = gotils.MustParseInt(m[2])
	part["s"] = gotils.MustParseInt(m[3])

	return part
}
