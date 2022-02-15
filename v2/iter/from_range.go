package iter

import (
	"math"

	"github.com/henrylee2cn/ameda/v2/ops"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type rangeNext[T Integer] struct {
	start       T
	end         T
	nextValue   T
	rightClosed bool
	ended       bool
}

func FromRange[T Integer](start T, end T, rightClosed ...bool) Iterator[T] {
	is := false
	if len(rightClosed) > 0 {
		is = rightClosed[0]
	}
	var next = rangeNext[T]{
		start:       start,
		end:         end,
		nextValue:   start,
		rightClosed: is,
	}
	return New[T](&next)
}

func (v *rangeNext[T]) Next() ops.Option[T] {
	if v.ended {
		return ops.None[T]()
	}
	if v.nextValue < v.end {
		value := v.nextValue
		v.nextValue++
		return ops.Some(value)
	}
	if v.nextValue == v.end {
		if v.rightClosed {
			return ops.Some(v.end)
		}
	}
	v.ended = true
	return ops.None[T]()
}

func (v *rangeNext[T]) SizeHint() (int, ops.Option[int]) {
	size := v.end - v.start
	if size < 0 {
		return 0, ops.Some(0)
	}
	if size == 0 {
		if v.rightClosed {
			return 0, ops.Some(1)
		}
		return 0, ops.Some(0)
	}
	up := uint64(size)
	if v.rightClosed {
		up += 1
	}
	if up > math.MaxInt {
		return 0, ops.None[int]()
	}
	return 0, ops.Some(int(up))
}
