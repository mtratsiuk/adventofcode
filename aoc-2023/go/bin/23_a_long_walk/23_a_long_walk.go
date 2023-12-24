package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("23_a_long_walk")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	grid := strings.Split(in, "\n")
	width := len(grid[0])
	height := len(grid)
	start := Pos2d{1, 0}
	end := Pos2d{width - 2, height - 1}

	maxSteps := 0

	var run func(pos, lastPos Pos2d, steps int)
	run = func(pos, lastPos Pos2d, steps int) {
		if pos.IsOutOfBounds(width, height) ||
			grid[pos.y][pos.x] == '#' {
			return
		}

		if pos == end {
			maxSteps = max(maxSteps, steps)
			return
		}

		for _, m := range AllMoves90 {
			nextPos := pos.Move(m)

			if nextPos == lastPos ||
				(grid[pos.y][pos.x] != '.' && m != ArrowToMove2d[grid[pos.y][pos.x]]) {
				continue
			}

			run(nextPos, pos, steps+1)
		}
	}

	run(start, start, 0)

	return maxSteps
}

func solve2(in string) int {
	grid := strings.Split(in, "\n")
	width := len(grid[0])
	height := len(grid)
	start := Pos2d{1, 0}
	end := Pos2d{width - 2, height - 1}
	nodes := make(map[Pos2d]*Node, 0)
	maxSteps := 0

	var build func(pos, lastPos, lastNodePos Pos2d, steps int)
	build = func(pos, lastPos, lastNodePos Pos2d, steps int) {
		if pos.IsOutOfBounds(width, height) ||
			grid[pos.y][pos.x] == '#' {
			return
		}

		if pos == end {
			nodes[pos] = NewNode(pos)
			nodes[lastNodePos].edges.Add(Edge{lastNodePos, pos, steps})
			nodes[pos].edges.Add(Edge{pos, lastNodePos, steps})
			return
		}

		if grid[pos.y][pos.x] != '.' {
			if nodes[lastNodePos] == nil {
				nodes[lastNodePos] = NewNode(lastNodePos)
			}

			intersection := pos.Move(ArrowToMove2d[grid[pos.y][pos.x]])

			if nodes[intersection] == nil {
				nodes[intersection] = NewNode(intersection)
			}

			nodes[intersection].edges.Add(Edge{intersection, lastNodePos, steps + 1})
			nodes[lastNodePos].edges.Add(Edge{lastNodePos, intersection, steps + 1})

			for _, m := range AllMoves90 {
				neighbor := intersection.Move(m)

				if neighbor == pos {
					continue
				}

				if grid[neighbor.y][neighbor.x] == '.' || grid[neighbor.y][neighbor.x] == '#' {
					continue
				}

				if neighbor.Move(ArrowToMove2d[grid[neighbor.y][neighbor.x]]) == intersection {
					continue
				}

				build(
					neighbor.Move(ArrowToMove2d[grid[neighbor.y][neighbor.x]]),
					neighbor,
					intersection,
					2)
			}

			return
		}

		for _, m := range AllMoves90 {
			nextPos := pos.Move(m)

			if nextPos == lastPos ||
				(grid[pos.y][pos.x] != '.' && m != ArrowToMove2d[grid[pos.y][pos.x]]) {
				continue
			}

			build(nextPos, pos, lastNodePos, steps+1)
		}
	}

	build(start, start, start, 0)

	var run func(node *Node, steps int, visited []bool)
	run = func(node *Node, steps int, visited []bool) {
		if node.pos == end {
			maxSteps = max(maxSteps, steps)
			return
		}

		for _, v := range node.edges.Items() {
			key := node.pos.x*1_000 + node.pos.y
			if visited[key] {
				continue
			}

			visited[key] = true
			run(nodes[v.end], steps+v.len, visited)
			visited[key] = false
		}
	}

	run(nodes[start], 0, make([]bool, 1_000_000))

	return maxSteps
}

func NewNode(pos Pos2d) *Node {
	n := Node{}
	n.pos = pos
	n.edges = gotils.NewSet[Edge]()
	return &n
}

type Node struct {
	pos   Pos2d
	edges gotils.Set[Edge]
}

type Edge struct {
	start, end Pos2d
	len        int
}

type Pos2d struct {
	x, y int
}

func (p Pos2d) IsOutOfBounds(width, height int) bool {
	return p.x < 0 || p.x >= width || p.y < 0 || p.y >= height
}

func (p Pos2d) Move(m Move2d) Pos2d {
	return Pos2d{p.x + m.dx, p.y + m.dy}
}

type Move2d struct {
	dx, dy int
}

var (
	Move2dTop    Move2d = Move2d{0, -1}
	Move2dRight  Move2d = Move2d{1, 0}
	Move2dBottom Move2d = Move2d{0, 1}
	Move2dLeft   Move2d = Move2d{-1, 0}
)

var AllMoves90 = []Move2d{
	Move2dTop,
	Move2dRight,
	Move2dBottom,
	Move2dLeft,
}

var ArrowToMove2d = map[byte]Move2d{
	'>': Move2dRight,
	'v': Move2dBottom,
	'<': Move2dLeft,
	'^': Move2dTop,
}
