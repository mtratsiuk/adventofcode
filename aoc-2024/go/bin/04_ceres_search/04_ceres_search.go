package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("04_ceres_search")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	rows := strings.Split(in, "\n")
	size := 4
	sum := 0
	seen := make(map[string]struct{}, 0)

	for y := 0; y <= len(rows)-size; y += 1 {
		for x := 0; x <= len(rows[0])-size; x += 1 {
			window := Window{
				data: rows,
				lrx:  x,
				lry:  y,
				size: size,
			}

			for view := range window.All() {
				if view.value == "XMAS" || view.value == "SAMX" {
					if _, ok := seen[view.Id()]; !ok {
						sum += 1
						seen[view.Id()] = struct{}{}
					}
				}
			}
		}
	}

	return sum
}

func solve2(in string) int {
	rows := strings.Split(in, "\n")
	size := 3
	sum := 0

	for y := 0; y <= len(rows)-size; y += 1 {
		for x := 0; x <= len(rows[0])-size; x += 1 {
			window := Window{
				data: rows,
				lrx:  x,
				lry:  y,
				size: size,
			}

			if gotils.Every(
				slices.Collect(window.Diagonals()),
				func(v View) bool { return v.value == "MAS" || v.value == "SAM" },
			) {
				sum += 1
			}
		}
	}

	return sum
}

type Window struct {
	data     []string
	lrx, lry int
	size     int
}

type View struct {
	value    string
	lrx, lry int
	tag      string
}

func (v *View) Id() string {
	return fmt.Sprintf("%v:%v:%v", v.lrx, v.lry, v.tag)
}

func (w *Window) All() iter.Seq[View] {
	return func(yield func(View) bool) {
		for view := range gotils.ConcatIters(w.Columns(), w.Rows(), w.Diagonals()) {
			if !yield(view) {
				return
			}
		}
	}
}

func (w *Window) Rows() iter.Seq[View] {
	return func(yield func(View) bool) {
		for y := w.lry; y < w.lry+w.size; y += 1 {
			if !yield(View{w.data[y][w.lrx : w.lrx+w.size], w.lrx, y, "r"}) {
				return
			}
		}
	}
}

func (w *Window) Columns() iter.Seq[View] {
	return func(yield func(View) bool) {
		for x := w.lrx; x < w.lrx+w.size; x += 1 {
			col := make([]byte, 0)

			for y := w.lry; y < w.lry+w.size; y += 1 {
				col = append(col, w.data[y][x])
			}

			if !yield(View{string(col), x, w.lry, "c"}) {
				return
			}
		}
	}
}

func (w *Window) Diagonals() iter.Seq[View] {
	return func(yield func(View) bool) {
		fwd := make([]byte, 0)
		bwd := make([]byte, 0)

		for i := 0; i < w.size; i += 1 {
			fwd = append(fwd, w.data[w.lry+i][w.lrx+i])
			bwd = append(bwd, w.data[w.lry+i][w.lrx+w.size-i-1])
		}

		if !yield(View{string(fwd), w.lrx, w.lry, "df"}) || !yield(View{string(bwd), w.lrx + w.size - 1, w.lry, "db"}) {
			return
		}
	}
}
