package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSlice_Ints(t *testing.T) {
	r := NewIntSlice([]int{1, 2, 3, 4, 5}).Ints()
	assert.Equal(t, []int{1, 2, 3, 4, 5}, r)
}

func TestIntSlice_Concat(t *testing.T) {
	a := []int{1}
	b := []int{2}
	c := []int{3}
	r := IntSlice(a).Concat(b, c)
	assert.Equal(t, []int{1, 2, 3}, r)
}

func TestIntSlice_CopyWithin(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	IntSlice(slice).CopyWithin(0, 3, 4)
	assert.Equal(t, []int{4, 2, 3, 4, 5}, slice)
	IntSlice(slice).CopyWithin(1, -2)
	assert.Equal(t, []int{4, 4, 5, 4, 5}, slice)
}

func TestIntSlice_Every(t *testing.T) {
	slice := []int{1, 30, 39, 29, 10, 13}
	isBelowThreshold := IntSlice(slice).Every(func(curr IntSlice, k int, v int) bool {
		return v < 40
	})
	assert.Equal(t, true, isBelowThreshold)
}

func TestIntSlice_Fill(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	IntSlice(slice).Fill(9999, 2, 4)
	assert.Equal(t, []int{1, 2, 9999, 9999}, slice)
	IntSlice(slice).Fill(5, -1)
	assert.Equal(t, []int{1, 2, 9999, 5}, slice)
}

func TestIntSlice_Filter(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	result := IntSlice(slice).Filter(func(curr IntSlice, k int, v int) bool {
		return v > 303
	})
	assert.Equal(t, []int{304, 305, 306}, result)
}

func TestIntSlice_Find(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	k, v := IntSlice(slice).Find(func(curr IntSlice, k int, v int) bool {
		return v > 303
	})
	assert.Equal(t, 3, k)
	assert.Equal(t, 304, v)
}

func TestIntSlice_Includes(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	had := IntSlice(slice).Includes(302)
	assert.True(t, had)
	had = IntSlice(slice).Includes(302, 1)
	assert.True(t, had)
	had = IntSlice(slice).Includes(302, 2)
	assert.False(t, had)
}

func TestIntSlice_IndexOf(t *testing.T) {
	slice := []int{301, 302, 303, 304, 305, 306}
	idx := IntSlice(slice).IndexOf(302)
	assert.Equal(t, 1, idx)
	idx = IntSlice(slice).IndexOf(302, 1)
	assert.Equal(t, 1, idx)
	idx = IntSlice(slice).IndexOf(302, 10)
	assert.Equal(t, -1, idx)
}

func TestIntSlice_LastIndexOf(t *testing.T) {
	slice := []int{101, 102, 103, 101}
	idx := IntSlice(slice).LastIndexOf(101)
	assert.Equal(t, 3, idx)
	idx = IntSlice(slice).LastIndexOf(101, 1)
	assert.Equal(t, 3, idx)
	idx = IntSlice(slice).LastIndexOf(101, 10)
	assert.Equal(t, -1, idx)
	idx = IntSlice(slice).LastIndexOf(9999)
	assert.Equal(t, -1, idx)
}

func TestIntSlice_Map(t *testing.T) {
	slice := []int{101, 102, 103, 101}
	ret := IntSlice(slice).Map(func(curr IntSlice, k int, v int) int {
		return k + v
	})
	assert.Equal(t, []int{101 + 0, 102 + 1, 103 + 2, 101 + 3}, ret)
}

func TestIntSlice_Pop(t *testing.T) {
	slice := IntSlice([]int{201, 202})
	last, ok := slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, 202, last)
	last, ok = slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, 201, last)
	last, ok = slice.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, last)
}

func TestIntSlice_PushOnce(t *testing.T) {
	slice := IntSlice{1, 2, 3, 4}
	n := slice.PushOnce(1, 5, 6, 1, 5, 6)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, IntSlice{1, 2, 3, 4, 5, 6}, slice)
}

func TestIntSlice_Reduce(t *testing.T) {
	slice := IntSlice([]int{1, 2, 3, 4})
	reducer := slice.Reduce(func(curr IntSlice, k int, v int, accumulator int) int {
		return accumulator - v
	})
	assert.Equal(t, 1-2-3-4, reducer)
	reducer = slice.Reduce(func(curr IntSlice, k int, v int, accumulator int) int {
		return accumulator - v
	}, 100)
	assert.Equal(t, 100-1-2-3-4, reducer)
}

func TestIntSlice_ReduceRight(t *testing.T) {
	slice := IntSlice([]int{1, 2, 3, 4})
	reducer := slice.ReduceRight(func(curr IntSlice, k int, v int, accumulator int) int {
		return accumulator - v
	})
	assert.Equal(t, 4-3-2-1, reducer)
	reducer = slice.ReduceRight(func(curr IntSlice, k int, v int, accumulator int) int {
		return accumulator - v
	}, 100)
	assert.Equal(t, 100-4-3-2-1, reducer)
}

func TestIntSlice_Reverse(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	IntSlice(slice).Reverse()
	assert.Equal(t, []int{4, 3, 2, 1}, slice)
}

func TestIntSlice_Shift(t *testing.T) {
	slice := IntSlice([]int{1, 2})
	first, ok := slice.Shift()
	assert.True(t, ok)
	assert.Equal(t, 1, first)
	first, ok = slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, first)
	first, ok = slice.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, first)
}

func TestIntSlice_Slice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	sub := IntSlice(slice).Slice(3)
	assert.Equal(t, []int{4, 5}, sub)
	sub = IntSlice(slice).Slice(3, 4)
	assert.Equal(t, []int{4}, sub)
	sub = IntSlice(slice).Slice(1, -2)
	assert.Equal(t, []int{2, 3}, sub)
	sub[0] = 999
	assert.Equal(t, []int{999, 3}, sub)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice)
}

func TestIntSlice_Some(t *testing.T) {
	slice := []int{1, 30, 39, 29, 10, 13}
	even := IntSlice(slice).Some(func(curr IntSlice, k int, v int) bool {
		return v%2 == 0
	})
	assert.Equal(t, true, even)
}

func TestIntSlice_Sort(t *testing.T) {
	slice := []int{1, 3, 4, 2}
	IntSlice(slice).Sort()
	assert.Equal(t, []int{1, 2, 3, 4}, slice)
}

func TestIntSlice_Splice(t *testing.T) {
	slice := IntSlice{0, 1, 2, 3, 4}
	slice.Splice(0, 0, 1, 2)
	assert.Equal(t, IntSlice{1, 2, 0, 1, 2, 3, 4}, slice)

	slice = IntSlice{0, 1, 2, 3, 4}
	slice.Splice(10, 0, 1, 2)
	assert.Equal(t, IntSlice{0, 1, 2, 3, 4, 1, 2}, slice)

	slice = IntSlice{0, 1, 2, 3, 4}
	slice.Splice(1, 0, 1, 2)
	assert.Equal(t, IntSlice{0, 1, 2, 1, 2, 3, 4}, slice)

	slice = IntSlice{0, 1, 2, 3, 4}
	slice.Splice(1, 2, 1, 2)
	assert.Equal(t, IntSlice{0, 1, 2, 3, 4}, slice)

	slice = IntSlice{0, 1, 2, 3, 4}
	slice.Splice(1, 10, 1, 2)
	assert.Equal(t, IntSlice{0, 1, 2}, slice)
}

func TestIntSlice_Unshift(t *testing.T) {
	slice := IntSlice{0, 1, 2, 3, 4}
	n := slice.Unshift(1, 2)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, IntSlice{1, 2, 0, 1, 2, 3, 4}, slice)
}

func TestIntSlice_UnshiftOnce(t *testing.T) {
	slice := IntSlice{1, 2, 3, 4}
	n := slice.UnshiftOnce(-1, 0, -1, 0, 1, 1)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, IntSlice{-1, 0, 1, 2, 3, 4}, slice)
}

func TestIntSlice_Distinct(t *testing.T) {
	slice := IntSlice{-1, 0, -1, 0, 1, 1}
	n := slice.Distinct()
	assert.Equal(t, len(slice), n)
	assert.Equal(t, IntSlice{-1, 0, 1}, slice)
}

func TestIntSlice_RemoveOne(t *testing.T) {
	slice := IntSlice{-1, 0, -1, 0, 1, 1}
	n := slice.RemoveOne(0)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, IntSlice{-1, -1, 0, 1, 1}, slice)
}

func TestIntSlice_RemoveEvery(t *testing.T) {
	slice := IntSlice{-1, 0, -1, 0, 1, 1}
	n := slice.RemoveEvery(0)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, IntSlice{-1, -1, 1, 1}, slice)
}
