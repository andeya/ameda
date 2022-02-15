package iter

import (
	"math"

	"github.com/henrylee2cn/ameda/v2/ops"
)

type chanNext[T comparable] struct {
	c <-chan T
}

func FromChan[T comparable](c <-chan T) Iterator[T] {
	return New[T](&chanNext[T]{c: c})
}

func (v *chanNext[T]) Next() ops.Option[T] {
	var x, ok = <-v.c
	if ok {
		return ops.Some(x)
	}
	return ops.None[T]()
}

func (v *chanNext[T]) SizeHint() (int, ops.Option[int]) {
	up := cap(v.c)
	if up > math.MaxInt {
		return 0, ops.None[int]()
	}
	return 0, ops.Some(up)
}
