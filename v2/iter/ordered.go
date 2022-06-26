package iter

import (
	"github.com/andeya/ameda/v2/ops"
	"github.com/andeya/ameda/v2/ord"
)

type orderedIterator[T ord.Ord] struct {
	baseIterator[T]
}

func newOrderedIterator[T ord.Ord](next Next[T]) *orderedIterator[T] {
	return &orderedIterator[T]{
		baseIterator: baseIterator[T]{next: next},
	}
}

func (iter *orderedIterator[T]) Max() ops.Option[T] {
	return iter.MaxBy(ord.Compare[T])
}

func (iter *orderedIterator[T]) Min() ops.Option[T] {
	return iter.MinBy(ord.Compare[T])
}

func (iter *orderedIterator[T]) MaxByKey(f func(x T) float64) ops.Option[T] {
	panic("")
	// type B struct {
	// 	p float64
	// 	t T
	// }
	// key := func(f func(x T) float64) func(T) B {
	// 	return func(x T) B {
	// 		return B{p: f(x), t: x}
	// 	}
	// }
	//
	// compare := func(x B, y B) ord.Ordering {
	// 	return ord.Compare(x.p, y.p)
	// }
	// iter2 := ToMap[T, B](iter, key(f))
	// r := iter2.Unwrap().MaxBy(compare)
	// if r.IsNone() {
	// 	return ops.None[T]()
	// }
	// return ops.Some(r.Some[B]())
}

func (iter *orderedIterator[T]) MaxBy(cmp func(x, y T) ord.Ordering) ops.Option[T] {
	return ops.None[T]()
}

func (iter *orderedIterator[T]) MinBy(cmp func(x, y T) ord.Ordering) ops.Option[T] {
	return ops.None[T]()
}
