package ameda

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsConcat(t *testing.T) {
	a := []string{"a"}
	b := []string{"b"}
	c := []string{"c"}
	r := StringsConcat(a, b, c)
	assert.Equal(t, []string{"a", "b", "c"}, r)
}

func TestStringsCopyWithin(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	StringsCopyWithin(slice, 0, 3, 4)
	assert.Equal(t, []string{"d", "b", "c", "d", "e"}, slice)
	StringsCopyWithin(slice, 1, -2)
	assert.Equal(t, []string{"d", "d", "e", "d", "e"}, slice)
}

func TestStringsEvery(t *testing.T) {
	slice := []string{"1", "30", "39", "29", "10", "13"}
	isBelowThreshold := StringsEvery(slice, func(s []string, k int, v string) bool {
		i, _ := strconv.Atoi(v)
		return i < 40
	})
	assert.Equal(t, true, isBelowThreshold)
}

func TestStringsFill(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	StringsFill(slice, "?", 2, 4)
	assert.Equal(t, []string{"a", "b", "?", "?"}, slice)
	StringsFill(slice, "e", -1)
	assert.Equal(t, []string{"a", "b", "?", "e"}, slice)
}

func TestStringsFilter(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	result := StringsFilter(slice, func(s []string, k int, v string) bool {
		return len(v) > 6
	})
	assert.Equal(t, []string{"exuberant", "destruction", "present"}, result)
}

func TestStringsFind(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	k, v := StringsFind(slice, func(s []string, k int, v string) bool {
		return len(v) > 6
	})
	assert.Equal(t, 3, k)
	assert.Equal(t, "exuberant", v)
}

func TestStringsIncludes(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	had := StringsIncludes(slice, "limit")
	assert.True(t, had)
	had = StringsIncludes(slice, "limit", 1)
	assert.True(t, had)
	had = StringsIncludes(slice, "limit", 2)
	assert.False(t, had)
}

func TestStringsIndexOf(t *testing.T) {
	slice := []string{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	idx := StringsIndexOf(slice, "limit")
	assert.Equal(t, 1, idx)
	idx = StringsIndexOf(slice, "limit", 1)
	assert.Equal(t, 1, idx)
	idx = StringsIndexOf(slice, "limit", 10)
	assert.Equal(t, -1, idx)
}

func TestStringsLastIndexOf(t *testing.T) {
	slice := []string{"Dodo", "Tiger", "Penguin", "Dodo"}
	idx := StringsLastIndexOf(slice, "Dodo")
	assert.Equal(t, 3, idx)
	idx = StringsLastIndexOf(slice, "Dodo", 1)
	assert.Equal(t, 3, idx)
	idx = StringsLastIndexOf(slice, "Dodo", 10)
	assert.Equal(t, -1, idx)
	idx = StringsLastIndexOf(slice, "?")
	assert.Equal(t, -1, idx)
}

func TestStringsMap(t *testing.T) {
	slice := []string{"Dodo", "Tiger", "Penguin", "Dodo"}
	ret := StringsMap(slice, func(s []string, k int, v string) string {
		return strconv.Itoa(k+1) + ":" + v
	})
	assert.Equal(t, []string{"1:Dodo", "2:Tiger", "3:Penguin", "4:Dodo"}, ret)
}

func TestStringsPop(t *testing.T) {
	slice := []string{"kale", "tomato"}
	last, ok := StringsPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, "tomato", last)
	last, ok = StringsPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, "kale", last)
	last, ok = StringsPop(&slice)
	assert.False(t, ok)
	assert.Equal(t, "", last)
}

func TestStringsPushOnce(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	n := StringsPushOnce(&slice, "1", "5", "6", "1", "5", "6")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, slice)
}

func TestStringsReduce(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	reducer := StringsReduce(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	})
	assert.Equal(t, "1+2+3+4", reducer)
	reducer = StringsReduce(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	}, "100")
	assert.Equal(t, "100+1+2+3+4", reducer)
}

func TestStringsReduceRight(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	reducer := StringsReduceRight(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	})
	assert.Equal(t, "4+3+2+1", reducer)
	reducer = StringsReduceRight(slice, func(s []string, k int, v string, accumulator string) string {
		return accumulator + "+" + v
	}, "100")
	assert.Equal(t, "100+4+3+2+1", reducer)
}

func TestStringsReverse(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	StringsReverse(slice)
	assert.Equal(t, []string{"4", "3", "2", "1"}, slice)
}

func TestStringsShift(t *testing.T) {
	slice := []string{"kale", "tomato"}
	first, ok := StringsShift(&slice)
	assert.True(t, ok)
	assert.Equal(t, "kale", first)
	first, ok = StringsPop(&slice)
	assert.True(t, ok)
	assert.Equal(t, "tomato", first)
	first, ok = StringsPop(&slice)
	assert.False(t, ok)
	assert.Equal(t, "", first)
}

func TestStringsSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	sub := StringsSlice(slice, 3)
	assert.Equal(t, []string{"d", "e"}, sub)
	sub = StringsSlice(slice, 3, 4)
	assert.Equal(t, []string{"d"}, sub)
	sub = StringsSlice(slice, 1, -2)
	assert.Equal(t, []string{"b", "c"}, sub)
	sub[0] = "x"
	assert.Equal(t, []string{"x", "c"}, sub)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, slice)
}

func TestStringsSome(t *testing.T) {
	slice := []string{"1", "30", "39", "29", "10", "13"}
	even := StringsSome(slice, func(s []string, k int, v string) bool {
		i, _ := strconv.Atoi(v)
		return i%2 == 0
	})
	assert.Equal(t, true, even)
}

func TestStringsSplice(t *testing.T) {
	slice := []string{"0", "1", "2", "3", "4"}
	StringsSplice(&slice, 0, 0, "a", "b")
	assert.Equal(t, []string{"a", "b", "0", "1", "2", "3", "4"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	StringsSplice(&slice, 10, 0, "a", "b")
	assert.Equal(t, []string{"0", "1", "2", "3", "4", "a", "b"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	StringsSplice(&slice, 1, 0, "a", "b")
	assert.Equal(t, []string{"0", "a", "b", "1", "2", "3", "4"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	StringsSplice(&slice, 1, 2, "a", "b")
	assert.Equal(t, []string{"0", "a", "b", "3", "4"}, slice)

	slice = []string{"0", "1", "2", "3", "4"}
	StringsSplice(&slice, 1, 10, "a", "b")
	assert.Equal(t, []string{"0", "a", "b"}, slice)
}

func TestStringsUnshift(t *testing.T) {
	slice := []string{"0", "1", "2", "3", "4"}
	n := StringsUnshift(&slice, "a", "b")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"a", "b", "0", "1", "2", "3", "4"}, slice)
}

func TestStringsUnshiftOnce(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	n := StringsUnshiftOnce(&slice, "-1", "0", "-1", "0", "1", "1")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "0", "1", "2", "3", "4"}, slice)
}

func TestStringsDistinct(t *testing.T) {
	slice := []string{"-1", "0", "-1", "0", "1", "1"}
	n := StringsDistinct(&slice)
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "0", "1"}, slice)
}

func TestStringsRemoveFirst(t *testing.T) {
	slice := []string{"-1", "0", "-1", "0", "1", "1"}
	n := StringsRemoveFirst(&slice, "0")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "-1", "0", "1", "1"}, slice)
}

func TestStringsRemoveEvery(t *testing.T) {
	slice := []string{"-1", "0", "-1", "0", "1", "1"}
	n := StringsRemoveEvery(&slice, "0")
	assert.Equal(t, len(slice), n)
	assert.Equal(t, []string{"-1", "-1", "1", "1"}, slice)
}
