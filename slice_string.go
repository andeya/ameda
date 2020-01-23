package ameda

import (
	"sort"
	"strconv"
	"strings"
)

// StringSlice string slice object
type StringSlice []string

// NewStringSlice creates a StringSlice object.
func NewStringSlice(a []string) *StringSlice {
	s := StringSlice(a)
	return &s
}

// Strings converts to []string.
func (s StringSlice) Strings() []string {
	return []string(s)
}

// Copy creates a copy of the string slice.
func (s StringSlice) Copy() []string {
	b := make([]string, len(s))
	copy(b, s)
	return b
}

// Bools converts string slice to bool slice.
func (s StringSlice) Bools(emptyAsZero ...bool) ([]bool, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]bool, len(s))
	for k, v := range s {
		i, err := strconv.ParseBool(v)
		if err != nil && strict {
			return nil, err
		}
		r[k] = i
	}
	return r, nil
}

// Float32s converts string slice to float32 slice.
func (s StringSlice) Float32s(emptyAsZero ...bool) ([]float32, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]float32, len(s))
	for k, v := range s {
		i, err := strconv.ParseFloat(v, 32)
		if err != nil && strict {
			return nil, err
		}
		r[k] = float32(i)
	}
	return r, nil
}

// Float64s converts string slice to float64 slice.
func (s StringSlice) Float64s(emptyAsZero ...bool) ([]float64, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]float64, len(s))
	for k, v := range s {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil && strict {
			return nil, err
		}
		r[k] = i
	}
	return r, nil
}

// Ints converts string slice to int slice.
func (s StringSlice) Ints(emptyAsZero ...bool) ([]int, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]int, len(s))
	for k, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil && strict {
			return nil, err
		}
		r[k] = i
	}
	return r, nil
}

// Int64s converts string slice to int64 slice.
func (s StringSlice) Int64s(emptyAsZero ...bool) ([]int64, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]int64, len(s))
	for k, v := range s {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil && strict {
			return nil, err
		}
		r[k] = i
	}
	return r, nil
}

// Int32s converts string slice to int32 slice.
func (s StringSlice) Int32s(emptyAsZero ...bool) ([]int32, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]int32, len(s))
	for k, v := range s {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil && strict {
			return nil, err
		}
		r[k] = int32(i)
	}
	return r, nil
}

// Int16s converts string slice to int16 slice.
func (s StringSlice) Int16s(emptyAsZero ...bool) ([]int16, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]int16, len(s))
	for k, v := range s {
		i, err := strconv.ParseInt(v, 10, 16)
		if err != nil && strict {
			return nil, err
		}
		r[k] = int16(i)
	}
	return r, nil
}

// Int8s converts string slice to int8 slice.
func (s StringSlice) Int8s(emptyAsZero ...bool) ([]int8, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]int8, len(s))
	for k, v := range s {
		i, err := strconv.ParseInt(v, 10, 8)
		if err != nil && strict {
			return nil, err
		}
		r[k] = int8(i)
	}
	return r, nil
}

// Uint8s converts string slice to uint8 slice.
func (s StringSlice) Uint8s(emptyAsZero ...bool) ([]uint8, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]uint8, len(s))
	for k, v := range s {
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil && strict {
			return nil, err
		}
		r[k] = uint8(i)
	}
	return r, nil
}

// Uint16s converts string slice to uint16 slice.
func (s StringSlice) Uint16s(emptyAsZero ...bool) ([]uint16, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]uint16, len(s))
	for k, v := range s {
		i, err := strconv.ParseUint(v, 10, 16)
		if err != nil && strict {
			return nil, err
		}
		r[k] = uint16(i)
	}
	return r, nil
}

// Uint32s converts string slice to uint32 slice.
func (s StringSlice) Uint32s(emptyAsZero ...bool) ([]uint32, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]uint32, len(s))
	for k, v := range s {
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil && strict {
			return nil, err
		}
		r[k] = uint32(i)
	}
	return r, nil
}

// Uint64s converts string slice to uint64 slice.
func (s StringSlice) Uint64s(emptyAsZero ...bool) ([]uint64, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]uint64, len(s))
	for k, v := range s {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil && strict {
			return nil, err
		}
		r[k] = uint64(i)
	}
	return r, nil
}

// Uints converts string slice to uint slice.
func (s StringSlice) Uints(emptyAsZero ...bool) ([]uint, error) {
	strict := !isEmptyAsZero(emptyAsZero)
	r := make([]uint, len(s))
	for k, v := range s {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil && strict {
			return nil, err
		}
		r[k] = uint(i)
	}
	return r, nil
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (s StringSlice) Concat(a ...[]string) []string {
	totalLen := len(s)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]string, totalLen)
	n := copy(ret, s)
	dst := ret[n:]
	for _, v := range a {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// CopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func (s StringSlice) CopyWithin(target, start int, end ...int) {
	target = s.fixIndex(target, true)
	if target == len(s) {
		return
	}
	sub := s.Slice(start, end...)
	for i, v := range sub {
		s[target+i] = v
	}
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func (s StringSlice) Every(fn func(curr StringSlice, k int, v string) bool) bool {
	for k, v := range s {
		if !fn(s, k, v) {
			return false
		}
	}
	return true
}

// Fill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func (s StringSlice) Fill(value string, start int, end ...int) {
	fixedStart, fixedEnd, ok := s.fixRange(start, end...)
	if !ok {
		return
	}
	for i := fixedStart; i < fixedEnd; i++ {
		s[i] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s StringSlice) Filter(fn func(curr StringSlice, k int, v string) bool) []string {
	ret := make([]string, 0)
	for k, v := range s {
		if fn(s, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func (s StringSlice) Find(fn func(curr StringSlice, k int, v string) bool) (k int, v string) {
	for k, v := range s {
		if fn(s, k, v) {
			return k, v
		}
	}
	return -1, ""
}

// Includes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (s StringSlice) Includes(valueToFind string, fromIndex ...int) bool {
	return s.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (s StringSlice) IndexOf(searchElement string, fromIndex ...int) int {
	idx := s.fromIndex(fromIndex...)
	for k, v := range s[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// Join concatenates the elements of s to create a single string. The separator string
// sep is placed between elements in the resulting string.
func (s StringSlice) Join(sep string) string {
	return strings.Join([]string(s), sep)
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (s StringSlice) LastIndexOf(searchElement string, fromIndex ...int) int {
	idx := s.fromIndex(fromIndex...)
	for i := len(s) - 1; i >= idx; i-- {
		if searchElement == s[i] {
			return i
		}
	}
	return -1
}

// Map creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func (s StringSlice) Map(fn func(curr StringSlice, k int, v string) string) []string {
	ret := make([]string, len(s))
	for k, v := range s {
		ret[k] = fn(s, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (s *StringSlice) Pop() (string, bool) {
	a := *s
	if len(a) == 0 {
		return "", false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*s = a[:len(a):len(a)]
	return last, true
}

// Push adds one or more elements to the end of an slice and returns the new length of the slice.
func (s *StringSlice) Push(element ...string) int {
	*s = append(*s, element...)
	return len(*s)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (s *StringSlice) PushOnce(element ...string) int {
	a := *s
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*s = a
	return len(a)
}

// Reduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func (s StringSlice) Reduce(fn func(curr StringSlice, k int, v string, accumulator string) string, initialValue ...string) string {
	if len(s) == 0 {
		return ""
	}
	start := 0
	acc := s[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for i := start; i < len(s); i++ {
		acc = fn(s, i, s[i], acc)
	}
	return acc
}

// ReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func (s StringSlice) ReduceRight(fn func(curr StringSlice, k int, v string, accumulator string) string, initialValue ...string) string {
	if len(s) == 0 {
		return ""
	}
	end := len(s) - 1
	acc := s[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for i := end; i >= 0; i-- {
		acc = fn(s, i, s[i], acc)
	}
	return acc
}

// Reverse reverses an slice in place.
func (s StringSlice) Reverse() {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// Shift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func (s *StringSlice) Shift() (string, bool) {
	a := *s
	if len(a) == 0 {
		return "", false
	}
	first := a[0]
	a = a[1:]
	*s = a[:len(a):len(a)]
	return first, true
}

// Slice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func (s StringSlice) Slice(begin int, end ...int) []string {
	fixedStart, fixedEnd, ok := s.fixRange(begin, end...)
	if !ok {
		return []string{}
	}
	return s[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (s StringSlice) Some(fn func(curr StringSlice, k int, v string) bool) bool {
	for k, v := range s {
		if fn(s, k, v) {
			return true
		}
	}
	return false
}

// Sort sorts the elements of an slice in place and returns the sorted slice.
func (s StringSlice) Sort() {
	sort.Strings([]string(s))
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (s *StringSlice) Splice(start, deleteCount int, items ...string) {
	a := *s
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := a.fixRange(start, start+1+deleteCount)
	deleteCount = end - start - 1
	for i := 0; i < len(items); i++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[i]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := a[start:].Copy()
			items = items[i:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*s = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*s = a[:len(a):len(a)]
}

// Unshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func (s *StringSlice) Unshift(element ...string) int {
	*s = append(element, *s...)
	return len(*s)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (s *StringSlice) UnshiftOnce(element ...string) int {
	a := *s
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[string]bool, len(element))
	r := make([]string, 0, len(a)+len(element))
L:
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		r = append(r, v)
	}
	r = append(r, a...)
	*s = r[:len(r):len(r)]
	return len(r)
}

// Distinct creates an new slice in place set that removes the same elements
// and returns the new length of the slice.
func (s *StringSlice) Distinct() int {
	a := (*s)[:0]
	m := make(map[string]bool, len(a))
	for _, v := range *s {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*s = a[:n:n]
	return n
}

// RemoveOne removes the first matched elements from the slice,
// and returns the new length of the slice.
func (s *StringSlice) RemoveOne(element ...string) int {
	a := *s
	m := make(map[string]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for i, vv := range a {
			if vv == v {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	n := len(a)
	*s = a[:n:n]
	return n
}

// RemoveAll removes all the elements from the slice,
// and returns the new length of the slice.
func (s *StringSlice) RemoveAll(element ...string) int {
	a := *s
	m := make(map[string]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for i, vv := range a {
			if vv == v {
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	n := len(a)
	*s = a[:n:n]
	return n
}

func (s StringSlice) fromIndex(fromIndex ...int) int {
	if len(fromIndex) > 0 {
		return s.fixIndex(fromIndex[0], true)
	}
	return 0
}

func (s StringSlice) fixRange(start int, end ...int) (fixedStart, fixedEnd int, ok bool) {
	fixedStart = s.fixIndex(start, true)
	if fixedStart == len(s) {
		return
	}
	fixedEnd = len(s)
	if len(end) > 0 {
		fixedEnd = s.fixIndex(end[0], true)
	}
	if fixedEnd-fixedStart <= 0 {
		return
	}
	ok = true
	return
}

func (s StringSlice) fixIndex(idx int, canLen bool) int {
	if idx < 0 {
		idx = len(s) + idx
		if idx < 0 {
			return 0
		}
		return idx
	}
	if idx >= len(s) {
		if canLen {
			return len(s)
		}
		return len(s) - 1
	}
	return idx
}

func isEmptyAsZero(emptyAsZero []bool) bool {
	return len(emptyAsZero) > 0 && emptyAsZero[0]
}
