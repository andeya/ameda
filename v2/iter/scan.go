package iter

import "github.com/henrylee2cn/ameda/v2/ops"

type Scan[T comparable] struct {
	iter         Iterator[T]
	initialState any
	f            func(*any, T) ops.Option[any]
}

func newScan[T comparable](iter Iterator[T], initialState any, f func(*any, T) ops.Option[any]) *Scan[T] {
	return &Scan[T]{iter: iter, initialState: initialState, f: f}
}
