package iter

import "github.com/andeya/ameda/v2/ops"

type MapWhile[T any] struct {
	iter Iterator[T]
	f    func(T) ops.Option[any]
}

func newMapWhile[T any](iter Iterator[T], f func(T) ops.Option[any]) *MapWhile[T] {
	return &MapWhile[T]{iter: iter, f: f}
}
