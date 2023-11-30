package gotils

import "fmt"

type Iterable[T any] interface {
	Next() (T, bool)
}

func Iter[T any](data []T) Iterable[T] {
	return &iterAll[T]{data, 0}
}

func Map[T, R any](it Iterable[T], f func(T) R) Iterable[R] {
	return &iterMap[T, R]{it, f}
}

func FlatMap[T, R any](it Iterable[T], f func(T) []R) Iterable[R] {
	return &iterFlatMap[T, R]{it, nil, f}
}

func Chunks[T any](it Iterable[T], n int) Iterable[[]T] {
	return &iterChunks[T]{it, n}
}

func Fold[T, R any](it Iterable[T], r R, f func(R, T) R) R {
	for n, done := it.Next(); !done; n, done = it.Next() {
		r = f(r, n)
	}

	return r
}

func Collect[T any](it Iterable[T]) []T {
	r := make([]T, 0)

	for n, done := it.Next(); !done; n, done = it.Next() {
		r = append(r, n)
	}

	return r
}

type iterAll[T any] struct {
	data []T
	i    int
}

func (it *iterAll[T]) Next() (next T, done bool) {
	if it.i < len(it.data) {
		next = it.data[it.i]
		it.i += 1
	} else {
		done = true
	}

	return
}

type iterMap[T, R any] struct {
	iter Iterable[T]
	f    func(T) R
}

func (it *iterMap[T, R]) Next() (R, bool) {
	n, done := it.iter.Next()

	if done {
		var r R
		return r, true
	}

	return it.f(n), false
}

type iterFlatMap[T, R any] struct {
	iter    Iterable[T]
	iterSub Iterable[R]
	f       func(T) []R
}

func (it *iterFlatMap[T, R]) Next() (R, bool) {
	if it.iterSub == nil {
		n, done := it.iter.Next()

		if done {
			var r R
			return r, true
		}

		mapped := it.f(n)
		it.iterSub = &iterAll[R]{mapped, 0}
		return it.iterSub.Next()
	}

	n, done := it.iterSub.Next()

	if done {
		it.iterSub = nil
		return it.Next()
	}

	return n, false
}

type iterChunks[T any] struct {
	iter Iterable[T]
	n    int
}

func (it *iterChunks[T]) Next() ([]T, bool) {
	chunk := make([]T, it.n)

	for i := 0; i < it.n; i += 1 {
		n, done := it.iter.Next()

		if done {
			if i == 0 {
				return chunk, true
			}

			panic(fmt.Sprintf("Chunks: not enough element to fill chunk: expected %v, iterator done at %v", it.n, i))
		}

		chunk[i] = n
	}

	return chunk, false
}
