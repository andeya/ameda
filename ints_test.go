package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsConcat(t *testing.T) {
	a := []int{1}
	b := []int{2}
	c := []int{3}
	r := IntsConcat(a, b, c)
	assert.Equal(t, []int{1, 2, 3}, r)
}

func TestIntsCopyWithin(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	IntsCopyWithin(slice, 0, 3, 4)
	assert.Equal(t, []int{4, 2, 3, 4, 5}, slice)
	IntsCopyWithin(slice, 1, -2)
	assert.Equal(t, []int{4, 4, 5, 4, 5}, slice)
}

func TestIntsEvery(t *testing.T) {
	slice := []int{1, 30, 39, 29, 10, 13}
	isBelowThreshold := IntsEvery(slice, func(i []int, k int, v int) bool {
		return v < 40
	})
	assert.Equal(t, true, isBelowThreshold)
}

func TestIntsFill(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	IntsFill(slice, 9999, 2, 4)
	assert.Equal(t, []int{1, 2, 9999, 9999}, slice)
	IntsFill(slice, 5, -1)
	assert.Equal(t, []int{1, 2, 9999, 5}, slice)
}

func TestIntsFilter(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	result := IntsFilter(slice, func(i []int, k int, v int) bool {
		return v > 303
	})
	assert.Equal(t, []int{304, 305, 306}, result)
}

func TestIntsFind(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	k, v := IntsFind(slice, func(i []int, k int, v int) bool {
		return v > 303
	})
	assert.Equal(t, 3, k)
	assert.Equal(t, 304, v)
}

func TestIntsIncludes(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	had := IntsIncludes(slice, 302)
	assert.True(t, had)
	had = IntsIncludes(slice, 302, 1)
	assert.True(t, had)
	had = IntsIncludes(slice, 302, 2)
	assert.False(t, had)
}

func TestIntsIndexOf(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	idx := IntsIndexOf(slice, 302)
	assert.Equal(t, 1, idx)
	idx = IntsIndexOf(slice, 302, 1)
	assert.Equal(t, 1, idx)
	idx = IntsIndexOf(slice, 302, 10)
	assert.Equal(t, -1, idx)
}

func TestIntsLastIndexOf(t *testing.T) {
	slice := []int{101, 102, 103, 101}
	idx := IntsLastIndexOf(slice, 101)
	assert.Equal(t, 3, idx)
	idx = IntsLastIndexOf(slice, 101, 1)
	assert.Equal(t, 3, idx)
	idx = IntsLastIndexOf(slice, 101, 10)
	assert.Equal(t, -1, idx)
	idx = IntsLastIndexOf(slice, 9999)
	assert.Equal(t, -1, idx)
}

func TestIntsMap(t *testing.T) {
	slice := []int{101, 102, 103, 101}
	ret := IntsMap(slice, func(i []int, k int, v int) int {
		return k + v
	})
	assert.Equal(t, []int{101 + 0, 102 + 1, 103 + 2, 101 + 3}, ret)
}

func TestIntsPop(t *testing.T) {
	slice := []int{201, 202}
	last, ok := IntsPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, 202, last)
	last, ok = IntsPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, 201, last)
	last, ok = IntsPop(&slice)
	assert.False(t, ok)
	assert.Equal(t, 0, last)
}

func TestIntsPushDistinct(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	slice = IntsPushDistinct(slice, 1, 5, 6, 1, 5, 6)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, slice)
}

func TestIntsReduce(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	reducer := IntsReduce(slice, func(i []int, k int, v int, accumulator int) int {
		return accumulator - v
	})
	assert.Equal(t, 1-2-3-4, reducer)
	reducer = IntsReduce(slice, func(i []int, k int, v int, accumulator int) int {
		return accumulator - v
	}, 100)
	assert.Equal(t, 100-1-2-3-4, reducer)
}

func TestIntsReduceRight(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	reducer := IntsReduceRight(slice, func(i []int, k int, v int, accumulator int) int {
		return accumulator - v
	})
	assert.Equal(t, 4-3-2-1, reducer)
	reducer = IntsReduceRight(slice, func(i []int, k int, v int, accumulator int) int {
		return accumulator - v
	}, 100)
	assert.Equal(t, 100-4-3-2-1, reducer)
}

func TestIntsReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	IntsReverse(slice)
	assert.Equal(t, []int{4, 3, 2, 1}, slice)
}

func TestIntsShift(t *testing.T) {
	slice := []int{1, 2}
	first, ok := IntsShift(&slice)
	assert.True(t, ok)
	assert.Equal(t, 1, first)
	first, ok = IntsPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, 2, first)
	first, ok = IntsPop(&slice)
	assert.False(t, ok)
	assert.Equal(t, 0, first)
}

func TestIntsSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	sub := IntsSlice(slice, 3)
	assert.Equal(t, []int{4, 5}, sub)
	sub = IntsSlice(slice, 3, 4)
	assert.Equal(t, []int{4}, sub)
	sub = IntsSlice(slice, 1, -2)
	assert.Equal(t, []int{2, 3}, sub)
	sub[0] = 999
	assert.Equal(t, []int{999, 3}, sub)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice)
}

func TestIntsSome(t *testing.T) {
	slice := []int{1, 30, 39, 29, 10, 13}
	even := IntsSome(slice, func(i []int, k int, v int) bool {
		return v%2 == 0
	})
	assert.Equal(t, true, even)
}

func TestIntsSplice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4}
	IntsSplice(&slice, 0, 0, 1, 2)
	assert.Equal(t, []int{1, 2, 0, 1, 2, 3, 4}, slice)

	slice = []int{0, 1, 2, 3, 4}
	IntsSplice(&slice, 10, 0, 1, 2)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 1, 2}, slice)

	slice = []int{0, 1, 2, 3, 4}
	IntsSplice(&slice, 1, 0, 1, 2)
	assert.Equal(t, []int{0, 1, 2, 1, 2, 3, 4}, slice)

	slice = []int{0, 1, 2, 3, 4}
	IntsSplice(&slice, 1, 2, 1, 2)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, slice)

	slice = []int{0, 1, 2, 3, 4}
	IntsSplice(&slice, 1, 10, 1, 2)
	assert.Equal(t, []int{0, 1, 2}, slice)
}

func TestIntsUnshift(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4}
	n := IntsUnshift(&slice, 1, 2)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []int{1, 2, 0, 1, 2, 3, 4}, slice)
}

func TestIntsUnshiftDistinct(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	n := IntsUnshiftDistinct(&slice, -1, 0, -1, 0, 1, 1)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []int{-1, 0, 1, 2, 3, 4}, slice)
}

func TestIntsDistinct(t *testing.T) {
	slice := []int{-1, 0, -1, 0, 1, 1}
	r := IntsDistinct(&slice, true)
	assert.Equal(t, len(slice), len(r))
	assert.Equal(t, []int{-1, 0, 1}, slice)
	assert.Equal(t, map[int]int{-1: 2, 0: 2, 1: 2}, r)
}

func TestIntsRemoveOne(t *testing.T) {
	slice := []int{-1, 0, -1, 0, 1, 1}
	n := IntsRemoveFirst(&slice, 0)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []int{-1, -1, 0, 1, 1}, slice)
}

func TestIntsRemoveEvery(t *testing.T) {
	slice := []int{-1, 0, -1, 0, 1, 1}
	n := IntsRemoveEvery(&slice, 0)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []int{-1, -1, 1, 1}, slice)
}
