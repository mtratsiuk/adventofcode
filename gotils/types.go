package gotils

import (
	"slices"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Set[T comparable] struct {
	data []T
}

func NewSet[T comparable]() Set[T] {
	s := Set[T]{make([]T, 0)}
	return s
}

func (s *Set[T]) Add(v T) {
	if !slices.Contains(s.data, v) {
		s.data = append(s.data, v)
	}
}

func (s *Set[T]) Has(v T) bool {
	return slices.Contains(s.data, v)
}

func (s *Set[T]) Items() []T {
	return s.data
}
