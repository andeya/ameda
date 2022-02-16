package iter

import "github.com/henrylee2cn/ameda/v2/ops"

func FromChars[T ~string](s T) Iterator[rune] {
	return FromVec[rune]([]rune(s))
}

type vecNext[T any] struct {
	nextIndex int
	vec       []T
}

func FromVec[T any](vec []T) Iterator[T] {
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
