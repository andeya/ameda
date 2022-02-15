package iter

import "github.com/henrylee2cn/ameda/v2/ops"

type vecNext[T comparable] struct {
	nextIndex int
	vec       []T
}

func FromVec[T comparable](vec []T) Iterator[T] {
	var next = vecNext[T]{
		nextIndex: 0,
		vec:       vec,
	}
	return New[T](&next)
}

func (v *vecNext[T]) Next() ops.Option[T] {
	idx := v.nextIndex
	if len(v.vec) >= idx {
		return ops.None[T]()
	}
	v.nextIndex++
	return ops.Some(v.vec[idx])
}

func (v *vecNext[T]) SizeHint() (int, ops.Option[int]) {
	size := len(v.vec)
	return size, ops.Some(size)
}
