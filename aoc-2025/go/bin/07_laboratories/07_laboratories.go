package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("07_laboratories")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	_, root := parse(in)
	count := make(map[*Node]int, 0)
	seen := gotils.NewSet[*Node]()

	var run func(node *Node) int
	run = func(node *Node) int {
		if node == nil || seen.Has(node) {
			return 0
		}

		if len(node.children) == 0 {
			seen.Add(node)
			return 1
		}

		if c, ok := count[node]; ok {
			return c
		}

		left := run(node.Left())
		right := run(node.Right())

		count[node] = 1 + left + right
		seen.Add(node)

		return count[node]
	}

	run(root.children[0])

	return count[root.children[0]]
}

func solve2(in string) int {
	_, root := parse(in)
	timelines := make(map[*Node]int, 0)

	var run func(node *Node) int
	run = func(node *Node) int {
		if node == nil {
			return 1
		}

		if t, ok := timelines[node]; ok {
			return t
		}

		left := run(node.Left())
		right := run(node.Right())

		timelines[node] = left + right

		return timelines[node]
	}

	run(root.children[0])

	return timelines[root.children[0]]
}

type Node struct {
	children []*Node
}

func (n *Node) Left() *Node {
	if len(n.children) > 0 {
		return n.children[0]
	}
	return nil
}

func (n *Node) Right() *Node {
	if len(n.children) > 1 {
		return n.children[1]
	}
	return nil
}

func NewNode() *Node {
	n := Node{}
	n.children = make([]*Node, 0)
	return &n
}

func parse(in string) ([]string, *Node) {
	grid := strings.Fields(in)
	w := len(grid[0])
	h := len(grid)

	posToNode := make(map[gotils.Pos2d]*Node, 0)
	root := NewNode()
	pos := gotils.NewPos2d(strings.Index(grid[0], "S"), 1)

	var run func(pos, origin gotils.Pos2d, prev *Node)
	run = func(pos, origin gotils.Pos2d, prev *Node) {
		if pos.IsOutOfBounds(w, h) {
			return
		}

		if grid[pos.Y][pos.X] == '.' {
			run(pos.Move(gotils.DirS), origin, prev)
			return
		}

		if grid[pos.Y][pos.X] == '^' {
			var node *Node
			if n, ok := posToNode[pos]; ok {
				node = n
			} else {
				node = NewNode()
			}

			if !slices.Contains(prev.children, node) {
				prev.children = append(prev.children, node)
			}

			posToNode[origin] = node
			posToNode[pos] = node

			left := pos.Move(gotils.DirW)
			if n, ok := posToNode[left]; ok {
				if !slices.Contains(node.children, n) {
					node.children = append(node.children, n)
				}
			} else {
				run(left, left, node)
			}

			right := pos.Move(gotils.DirE)
			if n, ok := posToNode[right]; ok {
				if !slices.Contains(node.children, n) {
					node.children = append(node.children, n)
				}
			} else {
				run(right, right, node)
			}

			return
		}

		panic("what is up")
	}

	run(pos, pos, root)

	return grid, root
}
