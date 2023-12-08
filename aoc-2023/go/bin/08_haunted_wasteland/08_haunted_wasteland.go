package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("08_haunted_wasteland")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	m := ParseMap(in)
	m.src = "AAA"
	m.dst = func(s string) bool { return s == "ZZZ" }

	m.Traverse(false)

	return m.current
}

func solve2(in string) int {
	m := ParseMap(in)
	m.dst = func(s string) bool { return strings.HasSuffix(s, "Z") }

	maps := make([]Map, 0)
	starts := make([]string, 0)
	steps := make([][]int, 0)

	for _, n := range m.nodes {
		if strings.HasSuffix(n.id, "A") {
			starts = append(starts, n.id)
		}
	}

	for _, s := range starts {
		newm := m
		newm.src = s
		newm.current = 0
		newm.stepsTillZ = make([]int, 0)
		maps = append(maps, newm)
	}

	for i := range maps {
		m := maps[i]

		seenZ := make(map[string]struct{}, 0)
		m.Traverse(true)

		for {
			m.Traverse(false)
			m.stepsTillZ = append(m.stepsTillZ, m.current)

			_, seen := seenZ[m.src]

			if seen {
				break
			}

			seenZ[m.src] = struct{}{}
			m.Traverse(true)
		}

		steps = append(steps, m.stepsTillZ)
	}

	minLcm := math.MaxInt

	for _, combs := range gotils.Combine(steps) {
		minLcm = min(minLcm, gotils.LcmAll(combs))
	}

	return minLcm
}

type Map struct {
	turns   string
	src     string
	dst     func(string) bool
	current int
	nodes   map[string]Node

	stepsTillZ []int
}

type Node struct {
	id    string
	left  string
	right string
}

func (m *Map) Traverse(once bool) {
	for {
		turn := m.turns[m.current%len(m.turns)]
		node := m.nodes[m.src]

		if !once && m.dst(node.id) {
			break
		}

		if turn == 'L' {
			m.src = m.nodes[node.left].id
		} else {
			m.src = m.nodes[node.right].id
		}

		m.current += 1

		if once {
			break
		}
	}
}

func ParseMap(in string) Map {
	m := Map{}
	m.nodes = make(map[string]Node)
	m.current = 0
	m.stepsTillZ = make([]int, 0)

	leftRight, nodes, _ := strings.Cut(in, "\n\n")
	nodeLines := strings.Split(nodes, "\n")

	m.turns = leftRight

	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	for _, line := range nodeLines {
		node := Node{}

		ids := re.FindStringSubmatch(line)

		node.id = ids[1]
		node.left = ids[2]
		node.right = ids[3]

		m.nodes[node.id] = node
	}

	return m
}
