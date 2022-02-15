package iter

import (
	"fmt"

	"github.com/henrylee2cn/ameda/v2/ops"
	"github.com/henrylee2cn/ameda/v2/result"
)

var _ Iterator[struct{}] = new(baseIterator[struct{}])

func New[T comparable](next Next[T]) Iterator[T] {
	return &baseIterator[T]{next: next}
}

type baseIterator[T comparable] struct {
	next Next[T]
}

func (iter *baseIterator[T]) IntoIter() Iterator[T] {
	return iter
}

func (iter *baseIterator[T]) Next() ops.Option[T] {
	return iter.next.Next()
}

func (iter *baseIterator[T]) SizeHint() (int, ops.Option[int]) {
	if sizeHint, ok := iter.next.(SizeHint); ok {
		return sizeHint.SizeHint()
	}
	return 0, ops.None[int]()
}

func (iter *baseIterator[T]) Count() int {
	return iter.Fold(0, func(count any, _ T) any { return count.(int) + 1 }).(int)
}

func (iter *baseIterator[T]) Last() ops.Option[T] {
	return iter.Fold(ops.None[T](), func(_ any, x T) any { return ops.Some(x) }).(ops.Option[T])
}

func (iter *baseIterator[T]) AdvanceBy(n int) result.Result[struct{}] {
	for i := 0; i < n; i++ {
		res := iter.Next().OkOr()
		if res.IsErr() {
			return result.Err[struct{}](fmt.Errorf("%d", i))
		}
	}
	return result.Ok[struct{}](struct{}{})
}

func (iter baseIterator[T]) Nth(n int) ops.Option[T] {
	var res = iter.AdvanceBy(n)
	if res.IsErr() {
		return ops.None[T]()
	}
	return iter.Next()
}

func (iter *baseIterator[T]) StepBy(step int) *StepBy[T] {
	return newStepBy[T](iter, step)
}

func (iter *baseIterator[T]) Chain(other IntoIterator[T]) *Chain[T] {
	return newChain[T](other.IntoIter())
}

func (iter *baseIterator[T]) Zip(other IntoIterator[T]) *Zip[T] {
	return newZip[T](other.IntoIter())
}

func (iter *baseIterator[T]) Intersperse(separator T) *Intersperse[T] {
	return newIntersperse[T](iter, separator)
}

func (iter *baseIterator[T]) IntersperseWith(separator func() T) *IntersperseWith[T] {
	return newIntersperseWith[T](iter, separator)
}

func (iter *baseIterator[T]) Map(f func(T) any) *Map[T] {
	return newMap[T](iter, f)
}

func (iter *baseIterator[T]) ForEach(f func(T)) {
	var call = func(f func(T)) func(any, T) any {
		return func(_ any, item T) any {
			f(item)
			return nil
		}
	}
	iter.Fold(nil, call(f))
}

func (iter *baseIterator[T]) Filter(f func(T) bool) *Filter[T] {
	return newFilter[T](iter, f)
}

func (iter *baseIterator[T]) FilterMap(f func(T) ops.Option[any]) *FilterMap[T] {
	return newFilterMap[T](iter, f)
}

func (iter *baseIterator[T]) Enumerate() *Enumerate[T] {
	return newEnumerate[T](iter)
}

func (iter *baseIterator[T]) Peekable() *Peekable[T] {
	return newPeekable[T](iter)
}

func (iter *baseIterator[T]) SkipWhile(f func(T) bool) *SkipWhile[T] {
	return newSkipWhile[T](iter, f)
}

func (iter *baseIterator[T]) TakeWhile(f func(T) bool) *TakeWhile[T] {
	return newTakeWhile[T](iter, f)
}

func (iter *baseIterator[T]) MapWhile(f func(T) ops.Option[any]) *MapWhile[T] {
	return newMapWhile[T](iter, f)
}

func (iter *baseIterator[T]) Skip(n int) *Skip[T] {
	return newSkip[T](iter, n)
}

func (iter *baseIterator[T]) Take(n int) *Take[T] {
	return newTake[T](iter, n)
}

func (iter *baseIterator[T]) Scan(initialState any, f func(initialState *any, item T) ops.Option[any]) *Scan[T] {
	return newScan[T](iter, initialState, f)
}

func (iter *baseIterator[T]) FlatMap(f func(T) IntoIterator[any]) *FlatMap[T] {
	return newFlatMap[T](iter, f)
}

func (iter *baseIterator[T]) Flatten() *Flatten[T, any] {
	return newFlatten[T, any](iter)
}

func (iter *baseIterator[T]) Fuse() *Fuse[T] {
	return newFuse[T](iter)
}

func (iter *baseIterator[T]) Collect() []T {
	size, maxOps := iter.SizeHint()
	if maxOps.IsSome() {
		size = maxOps.Some()
	}
	c := make([]T, 0, size)
	for {
		v := iter.Next()
		if v.IsNone() {
			break
		}
		c = append(c, v.Some())
	}
	return c
}

func (iter *baseIterator[T]) Partition(f func(T) bool) ([]T, []T) {
	return nil, nil
}
