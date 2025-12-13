package gotils

import (
	"iter"
	"maps"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	s := Set[T]{make(map[T]struct{}, 0)}
	return s
}

func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[T]) Items() iter.Seq[T] {
	return maps.Keys(s.data)
}

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() Queue[T] {
	q := Queue[T]{}
	q.data = make([]T, 0)
	return q
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) PopLeft() T {
	if len(q.data) == 0 {
		panic("PopLeft: queue is empty")
	}

	v := q.data[0]

	if len(q.data) > 1 {
		q.data = q.data[1:]
	} else {
		q.data = q.data[:0]
	}

	return v
}

func (q *Queue[T]) Pop() T {
	if len(q.data) == 0 {
		panic("Pop: queue is empty")
	}

	v := q.data[len(q.data)-1]
	q.data = q.data[0 : len(q.data)-1]

	return v
}

func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Items() []T {
	return q.data
}

type Pos2d struct {
	X int
	Y int
}

func NewPos2d(x, y int) Pos2d {
	return Pos2d{x, y}
}

func (p Pos2d) Move(d Pos2d) Pos2d {
	return Pos2d{p.X + d.X, p.Y + d.Y}
}

func (p Pos2d) IsOutOfBounds(width, height int) bool {
	return p.X < 0 || p.X >= width || p.Y < 0 || p.Y >= height
}

var (
	DirN  = Pos2d{0, -1}
	DirNE = Pos2d{1, -1}
	DirE  = Pos2d{1, 0}
	DirSE = Pos2d{1, 1}
	DirS  = Pos2d{0, 1}
	DirSW = Pos2d{-1, 1}
	DirW  = Pos2d{-1, 0}
	DirNW = Pos2d{-1, -1}
)

var Directions90 = []Pos2d{
	DirW,
	DirE,
	DirN,
	DirS,
}

var DirectionsAll = []Pos2d{
	DirN,
	DirNE,
	DirE,
	DirSE,
	DirS,
	DirSW,
	DirW,
	DirNW,
}
