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

func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[T]) Items() iter.Seq[T] {
	return maps.Keys(s.data)
}
