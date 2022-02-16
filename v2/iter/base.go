package iter

import (
	"github.com/henrylee2cn/ameda/v2/ops"
	"github.com/henrylee2cn/ameda/v2/result"
)

var _ Iterator[struct{}] = new(baseIterator[struct{}])

func New[T any](next Next[T]) Iterator[T] {
	return &baseIterator[T]{next: next}
}

type baseIterator[T any] struct {
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
			return result.Err[struct{}](i)
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
	_ = iter.Fold(nil, call(f))
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
	var extend = func(f func(T) bool, left, right *[]T) func(any, T) any {
		return func(_ any, x T) any {
			if f(x) {
				*left = append(*left, x)
			} else {
				*right = append(*right, x)
			}
			return nil
		}
	}
	minSize, _ := iter.SizeHint()
	left := make([]T, 0, minSize/2)
	right := make([]T, 0, minSize/2)
	_ = iter.Fold(nil, extend(f, &left, &right))
	return left, right
}

func (iter *baseIterator[T]) PartitionInPlace(f func(T) bool) int {
	panic("unimplemented!")
}

func (iter *baseIterator[T]) IsPartitioned(predicate func(T) bool) bool {
	// Either all items test `true`, or the first clause stops at `false`
	// and we check that there are no more `true` items after that.
	return iter.All(predicate) || !iter.Any(predicate)
}

func (iter *baseIterator[T]) TryFold(init any, f func(any, T) result.Result[any]) result.Result[any] {
	var accum = result.Ok(init)
	for {
		x := iter.Next()
		if x.IsNone() {
			break
		}
		accum = f(accum, x.Some())
		if accum.IsErr() {
			return accum
		}
	}
	return accum
}

func (iter *baseIterator[T]) TryForEach(f func(T) error) error {
	var call = func(f func(T) error) func(any, T) result.Result[any] {
		return func(_ any, x T) result.Result[any] {
			return result.Err[any](f(x))
		}
	}
	return iter.TryFold(nil, call(f)).Err()
}

func (iter *baseIterator[T]) Fold(init any, f func(any, T) any) any {
	var accum = init
	for {
		x := iter.Next()
		if x.IsNone() {
			break
		}
		accum = f(accum, x.Some())
	}
	return accum
}

func (iter *baseIterator[T]) Reduce(f func(accum T, item T) T) ops.Option[T] {
	var first = iter.Next()
	if first.IsNone() {
		return first
	}
	return ops.Some(iter.Fold(first, func(accum any, item T) any {
		return f(accum.(T), item)
	}).(T))
}

func (iter *baseIterator[T]) All(predicate func(T) bool) bool {
	var check = func(f func(T) bool) func(any, T) result.Result[any] {
		return func(_ any, x T) result.Result[any] {
			if f(x) {
				return result.Ok[any](nil)
			} else {
				return result.Err[any](nil)
			}
		}
	}
	return iter.TryFold(nil, check(predicate)).IsOk()
}

func (iter *baseIterator[T]) Any(predicate func(T) bool) bool {
	var check = func(f func(T) bool) func(any, T) result.Result[any] {
		return func(_ any, x T) result.Result[any] {
			if f(x) {
				return result.Err[any](nil)
			} else {
				return result.Ok[any](nil)
			}
		}
	}
	return iter.TryFold(nil, check(predicate)).IsErr()
}

func (iter *baseIterator[T]) Find(predicate func(T) bool) ops.Option[T] {
	var check = func(f func(T) bool) func(any, T) result.Result[any] {
		return func(_ any, x T) result.Result[any] {
			if f(x) {
				return result.Err[any](x)
			} else {
				return result.Ok[any](nil)
			}
		}
	}
	r := iter.TryFold(nil, check(predicate))
	if r.IsErr() {
		return ops.Some[T](r.ErrVal().(T))
	}
	return ops.None[T]()
}
