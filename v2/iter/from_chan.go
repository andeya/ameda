package iter

import (
	"math"

	"github.com/andeya/ameda/v2/ops"
)

type chanNext[T any] struct {
	c <-chan T
}

func FromChan[T any](c <-chan T) Iterator[T] {
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
