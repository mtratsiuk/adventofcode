package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("04_printing_department")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	for range accessible(grid(in)) {
		sum += 1
	}

	return sum
}

func solve2(in string) int {
	sum := 0
	g, w, h := grid(in)
	marked := make([]gotils.Pos2d, 0)

	for {
		for _, p := range marked {
			sum += 1
			g[p.Y][p.X] = '.'
		}

		marked = make([]gotils.Pos2d, 0)

		for p := range accessible(g, w, h) {
			marked = append(marked, p)
		}

		if len(marked) == 0 {
			break
		}
	}

	return sum
}

func accessible(g [][]byte, w, h int) iter.Seq[gotils.Pos2d] {
	roll := byte('@')
	threshold := 4

	return func(yield func(gotils.Pos2d) bool) {
		for y := 0; y < h; y += 1 {
			for x := 0; x < w; x += 1 {
				if g[y][x] != roll {
					continue
				}

				pos := gotils.NewPos2d(x, y)
				rolls := 0

				for _, dir := range gotils.DirectionsAll {
					n := pos.Move(dir)

					if !n.IsOutOfBounds(w, h) && g[n.Y][n.X] == roll {
						rolls += 1
					}
				}

				if rolls < threshold {
					if !yield(pos) {
						return
					}
				}
			}
		}
	}
}

func grid(in string) ([][]byte, int, int) {
	lines := strings.Fields(in)
	g := gotils.Collect(
		gotils.Map(
			gotils.Iter(lines),
			func(line string) []byte {
				return gotils.Collect(
					gotils.Map(gotils.Iter(strings.Split(line, "")), func(c string) byte { return c[0] }),
				)
			},
		),
	)

	return g, len(lines[0]), len(lines)
}
