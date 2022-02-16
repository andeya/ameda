package iter

import "github.com/henrylee2cn/ameda/v2/ops"

type FilterMap[T any] struct {
	iter Iterator[T]
	f    func(T) ops.Option[any]
}

func newFilterMap[T any](iter Iterator[T], f func(T) ops.Option[any]) *FilterMap[T] {
	return &FilterMap[T]{iter: iter, f: f}
}
