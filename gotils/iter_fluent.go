package gotils

import (
	"iter"
	"slices"
)

type iterFluent[T any] struct {
	seq iter.Seq[T]
}

func IterFluent[T any](v []T) iterFluent[T] {
	return iterFluent[T]{slices.Values(v)}
}

func (it iterFluent[T]) Map[R any](f func(T) R) iterFluent[R] {
	return iterFluent[R]{
		func(yield func(R) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

func (it iterFluent[T]) Fold[R any](r R, f func(R, T) R) R {
	for v := range it.seq {
		r = f(r, v)
	}

	return r
}
