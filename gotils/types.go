package gotils

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](items []T) Set[T] {
	s := Set[T]{make(map[T]struct{})}

	for _, v := range items {
		s.Add(v)
	}

	return s
}

func (s Set[T]) Add(v T) Set[T] {
	s.data[v] = struct{}{}
	return s
}

func (s Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s Set[T]) Items() []T {
	its := make([]T, 0)

	for v := range s.data {
		its = append(its, v)
	}

	return its
}

func (s Set[T]) Copy() Set[T] {
	return NewSet(s.Items())
}

func (s Set[T]) HashCode() string {
	var sb strings.Builder

	for v := range s.data {
		sb.WriteString(fmt.Sprintf("%v", v))
	}

	return sb.String()
}
