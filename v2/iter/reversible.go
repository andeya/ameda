package iter

import (
	"github.com/andeya/ameda/v2/ops"
	"github.com/andeya/ameda/v2/result"
)

type reversibleIterator[T any] struct {
	baseIterator[T]
}

func newReversibleIterator[T any](next ExactSizeNext[T]) *reversibleIterator[T] {
	return &reversibleIterator[T]{
		baseIterator: baseIterator[T]{next: next},
	}
}

func (iter *reversibleIterator[T]) Len() int {
	return iter.baseIterator.next.(ExactSizeNext[T]).Len()
}

func (iter *reversibleIterator[T]) RPosition(predicate func(T) bool) ops.Option[int] {
	var check = func(f func(T) bool) func(int, T) result.Result[int] {
		return func(i int, x T) result.Result[int] {
			i -= 1
			if f(x) {
				return result.Err[int](i)
			} else {
				return result.Ok[int](i)
			}
		}
	}
	r := TryRFold[T, int](iter, iter.Len(), check(predicate))
	if r.IsErr() {
		return ops.Some[int](r.ErrVal().(int))
	}
	return ops.None[int]()
}
