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

func solve2(in string) int {
	wsStr, _, _ := strings.Cut(in, "\n\n")
	ws := ParseWorkflows(strings.Split(wsStr, "\n"))
	graph := make(map[string][]Vert, 0)

	for from, rules := range ws {
		verts := make([]Vert, 0)
		apply := id

		for _, r := range rules {
			v, a := r.IntoVert(from, apply)
			apply = a
			verts = append(verts, v)
		}

		graph[from] = verts
	}

	sum := 0

	var run func(PartRanges, string)
	run = func(pr PartRanges, node string) {
		if node == "R" {
			return
		}

		if node == "A" {
			m := 1
			for _, r := range pr {
				m *= (r.right - r.left + 1)
			}
			sum += m
		}

		for _, v := range graph[node] {
			run(v.apply(pr), v.to)
		}
	}

	init := PartRanges{
		"x": Range{1, 4000},
		"m": Range{1, 4000},
		"a": Range{1, 4000},
		"s": Range{1, 4000},
	}

	run(init, "in")

	return sum
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

type PartRangeApplier func(PartRanges) PartRanges

func id(v PartRanges) PartRanges {
	return v.Copy()
}

func compose(f, g PartRangeApplier) PartRangeApplier {
	return func(pr PartRanges) PartRanges {
		return f(g(pr))
	}
}

type Vert struct {
	from  string
	to    string
	apply PartRangeApplier
}

type Rule struct {
	test func(Part) bool
	dest string
	key  string
	op   string
	val  int
}

func (r Rule) IsCond() bool {
	return r.op != ""
}

func (r Rule) IntoVert(from string, apply PartRangeApplier) (Vert, PartRangeApplier) {
	if !r.IsCond() {
		return Vert{from, r.dest, compose(id, apply)}, nil
	}

	if r.op == "<" {
		apply1 := func(pr PartRanges) PartRanges {
			cp := pr.Copy()
			cp[r.key] = cp[r.key].ApplyRight(r.val - 1)
			return cp
		}

		apply2 := func(pr PartRanges) PartRanges {
			cp := pr.Copy()
			cp[r.key] = cp[r.key].ApplyLeft(r.val)
			return cp
		}

		return Vert{from, r.dest, compose(apply1, apply)}, compose(apply2, apply)
	}

	apply1 := func(pr PartRanges) PartRanges {
		cp := pr.Copy()
		cp[r.key] = cp[r.key].ApplyLeft(r.val + 1)
		return cp
	}

	apply2 := func(pr PartRanges) PartRanges {
		cp := pr.Copy()
		cp[r.key] = cp[r.key].ApplyRight(r.val)
		return cp
	}

	return Vert{from, r.dest, compose(apply1, apply)}, compose(apply2, apply)
}

type Range struct {
	left, right int
}

func (r Range) ApplyLeft(left int) Range {
	return Range{max(r.left, left), r.right}
}

func (r Range) ApplyRight(right int) Range {
	return Range{r.left, min(r.right, right)}
}

type PartRanges map[string]Range

func (pr PartRanges) Copy() PartRanges {
	cp := make(PartRanges, len(pr))

	for k, v := range pr {
		cp[k] = v
	}

	return cp
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
				op := string(test[1])
				threshold := gotils.MustParseInt(test[2:])

				rule.dest = dest
				rule.key = key
				rule.op = op
				rule.val = threshold

				if op == "<" {
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
