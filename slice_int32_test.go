package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32Slice_Ints(t *testing.T) {
	r := NewInt32Slice([]int32{1, 2, 3, 4, 5}).Ints()
	assert.Equal(t, []int32{1, 2, 3, 4, 5}, r)
}

func TestInt32Slice_Concat(t *testing.T) {
	a := []int32{1}
	b := []int32{2}
	c := []int32{3}
	r := Int32Slice(a).Concat(b, c)
	assert.Equal(t, []int32{1, 2, 3}, r)
}

func TestInt32Slice_CopyWithin(t *testing.T) {
	slice := []int32{1, 2, 3, 4, 5}
	Int32Slice(slice).CopyWithin(0, 3, 4)
	assert.Equal(t, []int32{4, 2, 3, 4, 5}, slice)
	Int32Slice(slice).CopyWithin(1, -2)
	assert.Equal(t, []int32{4, 4, 5, 4, 5}, slice)
}

func TestInt32Slice_Every(t *testing.T) {
	slice := []int32{1, 30, 39, 29, 10, 13}
	isBelowThreshold := Int32Slice(slice).Every(func(curr Int32Slice, k int, v int32) bool {
		return v < 40
	})
	assert.Equal(t, true, isBelowThreshold)
}

func TestInt32Slice_Fill(t *testing.T) {
	slice := []int32{1, 2, 3, 4}
	Int32Slice(slice).Fill(9999, 2, 4)
	assert.Equal(t, []int32{1, 2, 9999, 9999}, slice)
	Int32Slice(slice).Fill(5, -1)
	assert.Equal(t, []int32{1, 2, 9999, 5}, slice)
}

func TestInt32Slice_Filter(t *testing.T) {
	slice := []int32{301, 302, 303, 304, 305, 306}
	result := Int32Slice(slice).Filter(func(curr Int32Slice, k int, v int32) bool {
		return v > 303
	})
	assert.Equal(t, []int32{304, 305, 306}, result)
}

func TestInt32Slice_Find(t *testing.T) {
	slice := []int32{301, 302, 303, 304, 305, 306}
	k, v := Int32Slice(slice).Find(func(curr Int32Slice, k int, v int32) bool {
		return v > 303
	})
	assert.Equal(t, 3, k)
	assert.Equal(t, int32(304), v)
}

func TestInt32Slice_Includes(t *testing.T) {
	slice := []int32{301, 302, 303, 304, 305, 306}
	had := Int32Slice(slice).Includes(302)
	assert.True(t, had)
	had = Int32Slice(slice).Includes(302, 1)
	assert.True(t, had)
	had = Int32Slice(slice).Includes(302, 2)
	assert.False(t, had)
}

func TestInt32Slice_IndexOf(t *testing.T) {
	slice := []int32{301, 302, 303, 304, 305, 306}
	idx := Int32Slice(slice).IndexOf(302)
	assert.Equal(t, 1, idx)
	idx = Int32Slice(slice).IndexOf(302, 1)
	assert.Equal(t, 1, idx)
	idx = Int32Slice(slice).IndexOf(302, 10)
	assert.Equal(t, -1, idx)
}

func TestInt32Slice_LastIndexOf(t *testing.T) {
	slice := []int32{101, 102, 103, 101}
	idx := Int32Slice(slice).LastIndexOf(101)
	assert.Equal(t, 3, idx)
	idx = Int32Slice(slice).LastIndexOf(101, 1)
	assert.Equal(t, 3, idx)
	idx = Int32Slice(slice).LastIndexOf(101, 10)
	assert.Equal(t, -1, idx)
	idx = Int32Slice(slice).LastIndexOf(9999)
	assert.Equal(t, -1, idx)
}

func TestInt32Slice_Map(t *testing.T) {
	slice := []int32{101, 102, 103, 101}
	ret := Int32Slice(slice).Map(func(curr Int32Slice, k int, v int32) int32 {
		return int32(k) + v
	})
	assert.Equal(t, []int32{101 + 0, 102 + 1, 103 + 2, 101 + 3}, ret)
}

func TestInt32Slice_Pop(t *testing.T) {
	slice := Int32Slice([]int32{201, 202})
	last, ok := slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, int32(202), last)
	last, ok = slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, int32(201), last)
	last, ok = slice.Pop()
	assert.False(t, ok)
	assert.Equal(t, int32(0), last)
}

func TestInt32Slice_PushOnce(t *testing.T) {
	slice := Int32Slice{1, 2, 3, 4}
	n := slice.PushOnce(1, 5, 6, 1, 5, 6)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, Int32Slice{1, 2, 3, 4, 5, 6}, slice)
}

func TestInt32Slice_Reduce(t *testing.T) {
	slice := Int32Slice([]int32{1, 2, 3, 4})
	reducer := slice.Reduce(func(curr Int32Slice, k int, v int32, accumulator int32) int32 {
		return accumulator - v
	})
	assert.Equal(t, int32(1-2-3-4), reducer)
	reducer = slice.Reduce(func(curr Int32Slice, k int, v int32, accumulator int32) int32 {
		return accumulator - v
	}, 100)
	assert.Equal(t, int32(100-1-2-3-4), reducer)
}

func TestInt32Slice_ReduceRight(t *testing.T) {
	slice := Int32Slice([]int32{1, 2, 3, 4})
	reducer := slice.ReduceRight(func(curr Int32Slice, k int, v int32, accumulator int32) int32 {
		return accumulator - v
	})
	assert.Equal(t, int32(4-3-2-1), reducer)
	reducer = slice.ReduceRight(func(curr Int32Slice, k int, v int32, accumulator int32) int32 {
		return accumulator - v
	}, 100)
	assert.Equal(t, int32(100-4-3-2-1), reducer)
}

func TestInt32Slice_Reverse(t *testing.T) {
	slice := []int32{1, 2, 3, 4}
	Int32Slice(slice).Reverse()
	assert.Equal(t, []int32{4, 3, 2, 1}, slice)
}

func TestInt32Slice_Shift(t *testing.T) {
	slice := Int32Slice([]int32{1, 2})
	first, ok := slice.Shift()
	assert.True(t, ok)
	assert.Equal(t, int32(1), first)
	first, ok = slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, int32(2), first)
	first, ok = slice.Pop()
	assert.False(t, ok)
	assert.Equal(t, int32(0), first)
}

func TestInt32Slice_Slice(t *testing.T) {
	slice := []int32{1, 2, 3, 4, 5}
	sub := Int32Slice(slice).Slice(3)
	assert.Equal(t, []int32{4, 5}, sub)
	sub = Int32Slice(slice).Slice(3, 4)
	assert.Equal(t, []int32{4}, sub)
	sub = Int32Slice(slice).Slice(1, -2)
	assert.Equal(t, []int32{2, 3}, sub)
	sub[0] = 999
	assert.Equal(t, []int32{999, 3}, sub)
	assert.Equal(t, []int32{1, 2, 3, 4, 5}, slice)
}

func TestInt32Slice_Some(t *testing.T) {
	slice := []int32{1, 30, 39, 29, 10, 13}
	even := Int32Slice(slice).Some(func(curr Int32Slice, k int, v int32) bool {
		return v%2 == 0
	})
	assert.Equal(t, true, even)
}

func TestInt32Slice_Sort(t *testing.T) {
	slice := []int32{1, 3, 4, 2}
	Int32Slice(slice).Sort()
	assert.Equal(t, []int32{1, 2, 3, 4}, slice)
}

func TestInt32Slice_Splice(t *testing.T) {
	slice := Int32Slice{0, 1, 2, 3, 4}
	slice.Splice(0, 0, 1, 2)
	assert.Equal(t, Int32Slice{1, 2, 0, 1, 2, 3, 4}, slice)

	slice = Int32Slice{0, 1, 2, 3, 4}
	slice.Splice(10, 0, 1, 2)
	assert.Equal(t, Int32Slice{0, 1, 2, 3, 4, 1, 2}, slice)

	slice = Int32Slice{0, 1, 2, 3, 4}
	slice.Splice(1, 0, 1, 2)
	assert.Equal(t, Int32Slice{0, 1, 2, 1, 2, 3, 4}, slice)

	slice = Int32Slice{0, 1, 2, 3, 4}
	slice.Splice(1, 2, 1, 2)
	assert.Equal(t, Int32Slice{0, 1, 2, 3, 4}, slice)

	slice = Int32Slice{0, 1, 2, 3, 4}
	slice.Splice(1, 10, 1, 2)
	assert.Equal(t, Int32Slice{0, 1, 2}, slice)
}

func TestInt32Slice_Unshift(t *testing.T) {
	slice := Int32Slice{0, 1, 2, 3, 4}
	n := slice.Unshift(1, 2)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, Int32Slice{1, 2, 0, 1, 2, 3, 4}, slice)
}

func TestInt32Slice_UnshiftOnce(t *testing.T) {
	slice := Int32Slice{1, 2, 3, 4}
	n := slice.UnshiftOnce(-1, 0, -1, 0, 1, 1)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, Int32Slice{-1, 0, 1, 2, 3, 4}, slice)
}

func TestInt32Slice_Distinct(t *testing.T) {
	slice := Int32Slice{-1, 0, -1, 0, 1, 1}
	n := slice.Distinct()
	assert.Equal(t, len(slice), n)
	assert.Equal(t, Int32Slice{-1, 0, 1}, slice)
}

func TestInt32Slice_RemoveOne(t *testing.T) {
	slice := Int32Slice{-1, 0, -1, 0, 1, 1}
	n := slice.RemoveOne(0)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, Int32Slice{-1, -1, 0, 1, 1}, slice)
}

func TestInt32Slice_RemoveAll(t *testing.T) {
	slice := Int32Slice{-1, 0, -1, 0, 1, 1}
	n := slice.RemoveAll(0)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, Int32Slice{-1, -1, 1, 1}, slice)
}
