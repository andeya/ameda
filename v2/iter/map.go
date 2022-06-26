package iter

import (
	"github.com/andeya/ameda/v2/ops"
	"github.com/andeya/ameda/v2/ord"
)

type Map[T any] struct {
	Iterator[T]
	f func(T) any
}

func newMap[T any](iter Iterator[T], f func(T) any) *Map[T] {
	return &Map[T]{Iterator: iter, f: f}
}

func OrderedIterTryFromMap[T any, B ord.Ord](m *Map[T]) ops.Option[OrderedIterator[B]] {
	panic("unimplemented")
}
