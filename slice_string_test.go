package ameda

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Strings(t *testing.T) {
	r := NewStringSlice([]string{"a", "b", "c", "d", "e"}).Strings()
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, r)
}

func TestStringSlice_Concat(t *testing.T) {
	a := []string{"a"}
	b := []string{"b"}
	c := []string{"c"}
	r := StringSlice(a).Concat(b, c)
	assert.Equal(t, []string{"a", "b", "c"}, r)
}

func TestStringSlice_CopyWithin(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	StringSlice(slice).CopyWithin(0, 3, 4)
	assert.Equal(t, []string{"d", "b", "c", "d", "e"}, slice)
	StringSlice(slice).CopyWithin(1, -2)
	assert.Equal(t, []string{"d", "d", "e", "d", "e"}, slice)
}

func TestStringSlice_Every(t *testing.T) {
	slice := []string{"1", "30", "39", "29", "10", "13"}
	isBelowThreshold := StringSlice(slice).Every(func(curr StringSlice, k int, v string) bool {
		i, _ := strconv.Atoi(v)
		return i < 40
	})
	assert.Equal(t, true, isBelowThreshold)
}

func TestStringSlice_Fill(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	StringSlice(slice).Fill("?", 2, 4)
	assert.Equal(t, []string{"a", "b", "?", "?"}, slice)
	StringSlice(slice).Fill("e", -1)
	assert.Equal(t, []string{"a", "b", "?", "e"}, slice)
}

func TestStringSlice_Filter(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	result := StringSlice(slice).Filter(func(curr StringSlice, k int, v string) bool {
		return len(v) > 6
	})
	assert.Equal(t, []string{"exuberant", "destruction", "present"}, result)
}

func TestStringSlice_Find(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	k, v := StringSlice(slice).Find(func(curr StringSlice, k int, v string) bool {
		return len(v) > 6
	})
	assert.Equal(t, 3, k)
	assert.Equal(t, "exuberant", v)
}

func TestStringSlice_Includes(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	had := StringSlice(slice).Includes("limit")
	assert.True(t, had)
	had = StringSlice(slice).Includes("limit", 1)
	assert.True(t, had)
	had = StringSlice(slice).Includes("limit", 2)
	assert.False(t, had)
}

func TestStringSlice_IndexOf(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	idx := StringSlice(slice).IndexOf("limit")
	assert.Equal(t, 1, idx)
	idx = StringSlice(slice).IndexOf("limit", 1)
	assert.Equal(t, 1, idx)
	idx = StringSlice(slice).IndexOf("limit", 10)
	assert.Equal(t, -1, idx)
}

func TestStringSlice_LastIndexOf(t *testing.T) {
	slice := []string{"Dodo", "Tiger", "Penguin", "Dodo"}
	idx := StringSlice(slice).LastIndexOf("Dodo")
	assert.Equal(t, 3, idx)
	idx = StringSlice(slice).LastIndexOf("Dodo", 1)
	assert.Equal(t, 3, idx)
	idx = StringSlice(slice).LastIndexOf("Dodo", 10)
	assert.Equal(t, -1, idx)
	idx = StringSlice(slice).LastIndexOf("?")
	assert.Equal(t, -1, idx)
}

func TestStringSlice_Map(t *testing.T) {
	slice := []string{"Dodo", "Tiger", "Penguin", "Dodo"}
	ret := StringSlice(slice).Map(func(curr StringSlice, k int, v string) string {
		return strconv.Itoa(k+1) + ":" + v
	})
	assert.Equal(t, []string{"1:Dodo", "2:Tiger", "3:Penguin", "4:Dodo"}, ret)
}

func TestStringSlice_Pop(t *testing.T) {
	slice := StringSlice([]string{"kale", "tomato"})
	last, ok := slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, "tomato", last)
	last, ok = slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, "kale", last)
	last, ok = slice.Pop()
	assert.False(t, ok)
	assert.Equal(t, "", last)
}

func TestStringSlice_PushOnce(t *testing.T) {
	slice := StringSlice{"1", "2", "3", "4"}
	n := slice.PushOnce("1", "5", "6", "1", "5", "6")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, StringSlice{"1", "2", "3", "4", "5", "6"}, slice)
}

func TestStringSlice_Reduce(t *testing.T) {
	slice := StringSlice([]string{"1", "2", "3", "4"})
	reducer := slice.Reduce(func(curr StringSlice, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	})
	assert.Equal(t, "1+2+3+4", reducer)
	reducer = slice.Reduce(func(curr StringSlice, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	}, "100")
	assert.Equal(t, "100+1+2+3+4", reducer)
}

func TestStringSlice_ReduceRight(t *testing.T) {
	slice := StringSlice([]string{"1", "2", "3", "4"})
	reducer := slice.ReduceRight(func(curr StringSlice, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	})
	assert.Equal(t, "4+3+2+1", reducer)
	reducer = slice.ReduceRight(func(curr StringSlice, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	}, "100")
	assert.Equal(t, "100+4+3+2+1", reducer)
}

func TestStringSlice_Reverse(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	StringSlice(slice).Reverse()
	assert.Equal(t, []string{"4", "3", "2", "1"}, slice)
}

func TestStringSlice_Shift(t *testing.T) {
	slice := StringSlice([]string{"kale", "tomato"})
	first, ok := slice.Shift()
	assert.True(t, ok)
	assert.Equal(t, "kale", first)
	first, ok = slice.Pop()
	assert.True(t, ok)
	assert.Equal(t, "tomato", first)
	first, ok = slice.Pop()
	assert.False(t, ok)
	assert.Equal(t, "", first)
}

func TestStringSlice_Slice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	sub := StringSlice(slice).Slice(3)
	assert.Equal(t, []string{"d", "e"}, sub)
	sub = StringSlice(slice).Slice(3, 4)
	assert.Equal(t, []string{"d"}, sub)
	sub = StringSlice(slice).Slice(1, -2)
	assert.Equal(t, []string{"b", "c"}, sub)
	sub[0] = "x"
	assert.Equal(t, []string{"x", "c"}, sub)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, slice)
}

func TestStringSlice_Some(t *testing.T) {
	slice := []string{"1", "30", "39", "29", "10", "13"}
	even := StringSlice(slice).Some(func(curr StringSlice, k int, v string) bool {
		i, _ := strconv.Atoi(v)
		return i%2 == 0
	})
	assert.Equal(t, true, even)
}

func TestStringSlice_Sort(t *testing.T) {
	slice := []string{"1", "3", "4", "2"}
	StringSlice(slice).Sort()
	assert.Equal(t, []string{"1", "2", "3", "4"}, slice)
}

func TestStringSlice_Splice(t *testing.T) {
	slice := StringSlice{"0", "1", "2", "3", "4"}
	slice.Splice(0, 0, "a", "b")
	assert.Equal(t, StringSlice{"a", "b", "0", "1", "2", "3", "4"}, slice)

	slice = StringSlice{"0", "1", "2", "3", "4"}
	slice.Splice(10, 0, "a", "b")
	assert.Equal(t, StringSlice{"0", "1", "2", "3", "4", "a", "b"}, slice)

	slice = StringSlice{"0", "1", "2", "3", "4"}
	slice.Splice(1, 0, "a", "b")
	assert.Equal(t, StringSlice{"0", "a", "b", "1", "2", "3", "4"}, slice)

	slice = StringSlice{"0", "1", "2", "3", "4"}
	slice.Splice(1, 2, "a", "b")
	assert.Equal(t, StringSlice{"0", "a", "b", "3", "4"}, slice)

	slice = StringSlice{"0", "1", "2", "3", "4"}
	slice.Splice(1, 10, "a", "b")
	assert.Equal(t, StringSlice{"0", "a", "b"}, slice)
}

func TestStringSlice_Unshift(t *testing.T) {
	slice := StringSlice{"0", "1", "2", "3", "4"}
	n := slice.Unshift("a", "b")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, StringSlice{"a", "b", "0", "1", "2", "3", "4"}, slice)
}

func TestStringSlice_UnshiftOnce(t *testing.T) {
	slice := StringSlice{"1", "2", "3", "4"}
	n := slice.UnshiftOnce("-1", "0", "-1", "0", "1", "1")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, StringSlice{"-1", "0", "1", "2", "3", "4"}, slice)
}

func TestStringSlice_Distinct(t *testing.T) {
	slice := StringSlice{"-1", "0", "-1", "0", "1", "1"}
	n := slice.Distinct()
	assert.Equal(t, len(slice), n)
	assert.Equal(t, StringSlice{"-1", "0", "1"}, slice)
}

func TestStringSlice_RemoveOne(t *testing.T) {
	slice := StringSlice{"-1", "0", "-1", "0", "1", "1"}
	n := slice.RemoveOne("0")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, StringSlice{"-1", "-1", "0", "1", "1"}, slice)
}

func TestStringSlice_RemoveAll(t *testing.T) {
	slice := StringSlice{"-1", "0", "-1", "0", "1", "1"}
	n := slice.RemoveAll("0")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, StringSlice{"-1", "-1", "1", "1"}, slice)
}
