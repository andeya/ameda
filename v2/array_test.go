package ameda

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVecConcat(t *testing.T) {
	a := []string{"a", "0"}
	b := []string{"b", "1"}
	c := []string{"c", "2"}
	r := VecConcat(a, b, c)
	assert.Equal(t, []string{"a", "0", "b", "1", "c", "2"}, r)
}

func TestVecIntersect(t *testing.T) {
	x := []string{"a", "b", "a", "b", "b", "a", "a"}
	y := []string{"a", "b", "c", "a", "b", "c", "b", "c", "c"}
	z := []string{"a", "b", "c", "d", "a", "b", "c", "d", "b", "c", "d", "c", "d", "d"}
	r := VecIntersect(x, y, z)
	assert.Equal(t, map[string]int{"a": 2, "b": 3}, r)
}

func TestVecCopyWithin(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	VecCopyWithin(slice, 0, 3, 4)
	assert.Equal(t, []string{"d", "b", "c", "d", "e"}, slice)
	VecCopyWithin(slice, 1, -2)
	assert.Equal(t, []string{"d", "d", "e", "d", "e"}, slice)
}

func TestVecEvery(t *testing.T) {
	slice := []string{"1", "30", "39", "29", "10", "13"}
	isBelowThreshold := VecEvery(slice, func(s []string, k int, v string) bool {
		i, _ := strconv.Atoi(v)
		return i < 40
	})
	assert.Equal(t, true, isBelowThreshold)
}

func TestVecFill(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	VecFill(slice, "?", 2, 4)
	assert.Equal(t, []string{"a", "b", "?", "?"}, slice)
	VecFill(slice, "e", -1)
	assert.Equal(t, []string{"a", "b", "?", "e"}, slice)
}

func TestVecFilter(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	result := VecFilter(slice, func(s []string, k int, v string) bool {
		return len(v) > 6
	})
	assert.Equal(t, []string{"exuberant", "destruction", "present"}, result)
}

func TestVecFind(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	k, v := VecFind(slice, func(s []string, k int, v string) bool {
		return len(v) > 6
	})
	assert.Equal(t, 3, k)
	assert.Equal(t, "exuberant", v)
}

func TestVecIncludes(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	had := VecIncludes(slice, "limit")
	assert.True(t, had)
	had = VecIncludes(slice, "limit", 1)
	assert.True(t, had)
	had = VecIncludes(slice, "limit", 2)
	assert.False(t, had)
}

func TestVecIndexOf(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	idx := VecIndexOf(slice, "limit")
	assert.Equal(t, 1, idx)
	idx = VecIndexOf(slice, "limit", 1)
	assert.Equal(t, 1, idx)
	idx = VecIndexOf(slice, "limit", 10)
	assert.Equal(t, -1, idx)
}

func TestVecLastIndexOf(t *testing.T) {
	slice := []string{"Dodo", "Tiger", "Penguin", "Dodo"}
	idx := VecLastIndexOf(slice, "Dodo")
	assert.Equal(t, 3, idx)
	idx = VecLastIndexOf(slice, "Dodo", 1)
	assert.Equal(t, 3, idx)
	idx = VecLastIndexOf(slice, "Dodo", 10)
	assert.Equal(t, -1, idx)
	idx = VecLastIndexOf(slice, "?")
	assert.Equal(t, -1, idx)
}

func TestVecMap(t *testing.T) {
	slice := []string{"Dodo", "Tiger", "Penguin", "Dodo"}
	ret := VecMap(slice, func(s []string, k int, v string) string {
		return strconv.Itoa(k+1) + ":" + v
	})
	assert.Equal(t, []string{"1:Dodo", "2:Tiger", "3:Penguin", "4:Dodo"}, ret)
}

func TestVecPop(t *testing.T) {
	slice := []string{"kale", "tomato"}
	last, ok := VecPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, "tomato", last)
	last, ok = VecPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, "kale", last)
	last, ok = VecPop(&slice)
	assert.False(t, ok)
	assert.Equal(t, "", last)
}

func TestVecPushDistinct(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	slice = VecPushDistinct(slice, "1", "5", "6", "1", "5", "6")
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, slice)
}

func TestVecReduce(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	reducer := VecReduce(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	})
	assert.Equal(t, "1+2+3+4", reducer)
	reducer = VecReduce(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	}, "100")
	assert.Equal(t, "100+1+2+3+4", reducer)
}

func TestVecReduceRight(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	reducer := VecReduceRight(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	})
	assert.Equal(t, "4+3+2+1", reducer)
	reducer = VecReduceRight(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	}, "100")
	assert.Equal(t, "100+4+3+2+1", reducer)
}

func TestVecReverse(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	VecReverse(slice)
	assert.Equal(t, []string{"4", "3", "2", "1"}, slice)
}

func TestVecShift(t *testing.T) {
	slice := []string{"kale", "tomato"}
	first, ok := VecShift(&slice)
	assert.True(t, ok)
	assert.Equal(t, "kale", first)
	first, ok = VecPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, "tomato", first)
	first, ok = VecPop(&slice)
	assert.False(t, ok)
	assert.Equal(t, "", first)
}

func TestVecSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	sub := VecSlice(slice, 3)
	assert.Equal(t, []string{"d", "e"}, sub)
	sub = VecSlice(slice, 3, 4)
	assert.Equal(t, []string{"d"}, sub)
	sub = VecSlice(slice, 1, -2)
	assert.Equal(t, []string{"b", "c"}, sub)
	sub[0] = "x"
	assert.Equal(t, []string{"x", "c"}, sub)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, slice)
}

func TestVecSome(t *testing.T) {
	slice := []string{"1", "30", "39", "29", "10", "13"}
	even := VecSome(slice, func(s []string, k int, v string) bool {
		i, _ := strconv.Atoi(v)
		return i%2 == 0
	})
	assert.Equal(t, true, even)
}

func TestVecSplice(t *testing.T) {
	slice := []string{"0", "1", "2", "3", "4"}
	VecSplice(&slice, 0, 0, "a", "b")
	assert.Equal(t, []string{"a", "b", "0", "1", "2", "3", "4"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	VecSplice(&slice, 10, 0, "a", "b")
	assert.Equal(t, []string{"0", "1", "2", "3", "4", "a", "b"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	VecSplice(&slice, 1, 0, "a", "b")
	assert.Equal(t, []string{"0", "a", "b", "1", "2", "3", "4"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	VecSplice(&slice, 1, 2, "a", "b")
	assert.Equal(t, []string{"0", "a", "b", "3", "4"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	VecSplice(&slice, 1, 10, "a", "b")
	assert.Equal(t, []string{"0", "a", "b"}, slice)
}

func TestVecUnshift(t *testing.T) {
	slice := []string{"0", "1", "2", "3", "4"}
	n := VecUnshift(&slice, "a", "b")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"a", "b", "0", "1", "2", "3", "4"}, slice)
}

func TestVecUnshiftDistinct(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	n := VecUnshiftDistinct(&slice, "-1", "0", "-1", "0", "1", "1")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "0", "1", "2", "3", "4"}, slice)
}

func TestVecDistinct(t *testing.T) {
	slice := []string{"-1", "0", "-1", "0", "1", "1"}
	distinctCount := VecDistinct(&slice, true)
	assert.Equal(t, len(slice), len(distinctCount))
	assert.Equal(t, []string{"-1", "0", "1"}, slice)
	assert.Equal(t, map[string]int{"-1": 2, "0": 2, "1": 2}, distinctCount)
}

func TestVecRemoveFirst(t *testing.T) {
	slice := []string{"-1", "0", "-1", "0", "1", "1"}
	n := VecRemoveFirst(&slice, "0")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "-1", "0", "1", "1"}, slice)
}

func TestVecRemoveEvery(t *testing.T) {
	slice := []string{"-1", "0", "-1", "0", "1", "1"}
	n := VecRemoveEvery(&slice, "0")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "-1", "1", "1"}, slice)
}

func TestStringSet(t *testing.T) {
	set1 := []string{"1", "2", "3", "6", "8"}
	set2 := []string{"2", "3", "5", "0"}
	set3 := []string{"2", "6", "7"}
	un := SetsUnion(set1, set2, set3)
	assert.Equal(t, []string{"1", "2", "3", "6", "8", "5", "0", "7"}, un)
	in := SetsIntersect(set1, set2, set3)
	assert.Equal(t, []string{"2"}, in)
	di := SetsDifference(set1, set2, set3)
	assert.Equal(t, []string{"1", "8"}, di)
}
