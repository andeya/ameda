package ameda

import (
	"sort"
)

// Int32Slice int32 slice object
type Int32Slice []int32

// NewInt32Slice creates an Int32Slice object.
func NewInt32Slice(a []int32) *Int32Slice {
	i := Int32Slice(a)
	return &i
}

// Int32sCopy creates a copy of the int32 slice.
func Int32sCopy(i []int32) []int32 {
	b := make([]int32, len(i))
	copy(b, i)
	return b
}

// Copy creates a copy of the int32 slice.
func (i Int32Slice) Copy() []int32 {
	return Int32sCopy(i)
}

// Int32sToInterfaces converts int32 slice to interface slice.
func Int32sToInterfaces(i []int32) []interface{} {
	r := make([]interface{}, len(i))
	for k, v := range i {
		r[k] = Int32ToInterface(v)
	}
	return r
}

// Interfaces converts int32 slice to interface slice.
func (i Int32Slice) Interfaces() []interface{} {
	return Int32sToInterfaces(i)
}

// Int32sToStrings converts int32 slice to string slice.
func Int32sToStrings(i []int32) []string {
	r := make([]string, len(i))
	for k, v := range i {
		r[k] = Int32ToString(v)
	}
	return r
}

// Strings converts int32 slice to string slice.
func (i Int32Slice) Strings() []string {
	return Int32sToStrings(i)
}

// Int32sToBools converts int32 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func Int32sToBools(i []int32) []bool {
	r := make([]bool, len(i))
	for k, v := range i {
		r[k] = Int32ToBool(v)
	}
	return r
}

// Bools converts int32 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func (i Int32Slice) Bools() []bool {
	return Int32sToBools(i)
}

// Int32sToFloat32s converts int32 slice to float32 slice.
func Int32sToFloat32s(i []int32) []float32 {
	r := make([]float32, len(i))
	for k, v := range i {
		r[k] = Int32ToFloat32(v)
	}
	return r
}

// Float32s converts int32 slice to float32 slice.
func (i Int32Slice) Float32s() []float32 {
	return Int32sToFloat32s(i)
}

// Int32sToFloat64s converts int32 slice to float64 slice.
func Int32sToFloat64s(i []int32) []float64 {
	r := make([]float64, len(i))
	for k, v := range i {
		r[k] = Int32ToFloat64(v)
	}
	return r
}

// Float64s converts int32 slice to float64 slice.
func (i Int32Slice) Float64s() []float64 {
	return Int32sToFloat64s(i)
}

// Int32sToInts converts int32 slice to int slice.
func Int32sToInts(i []int32) []int {
	r := make([]int, len(i))
	for k, v := range i {
		r[k] = Int32ToInt(v)
	}
	return r
}

// Ints converts int32 slice to int slice.
func (i Int32Slice) Ints() []int {
	return Int32sToInts(i)
}

// Int32sToInt8s converts int32 slice to int8 slice.
func Int32sToInt8s(i []int32) ([]int8, error) {
	var err error
	r := make([]int8, len(i))
	for k, v := range i {
		r[k], err = Int32ToInt8(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Int8s converts int32 slice to int8 slice.
func (i Int32Slice) Int8s() ([]int8, error) {
	return Int32sToInt8s(i)
}

// Int32sToInt16s converts int32 slice to int16 slice.
func Int32sToInt16s(i []int32) ([]int16, error) {
	var err error
	r := make([]int16, len(i))
	for k, v := range i {
		r[k], err = Int32ToInt16(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Int16s converts int32 slice to int16 slice.
func (i Int32Slice) Int16s() ([]int16, error) {
	return Int32sToInt16s(i)
}

// Int32s converts to []int32.
func (i Int32Slice) Int32s() []int32 {
	return []int32(i)
}

// Int32sToInt64s converts int32 slice to int64 slice.
func Int32sToInt64s(i []int32) []int64 {
	r := make([]int64, len(i))
	for k, v := range i {
		r[k] = Int32ToInt64(v)
	}
	return r
}

// Int64s converts int32 slice to int64 slice.
func (i Int32Slice) Int64s() []int64 {
	return Int32sToInt64s(i)
}

// Int32sToUints converts int32 slice to uint slice.
func Int32sToUints(i []int32) ([]uint, error) {
	var err error
	r := make([]uint, len(i))
	for k, v := range i {
		r[k], err = Int32ToUint(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uints converts int32 slice to uint slice.
func (i Int32Slice) Uints() ([]uint, error) {
	return Int32sToUints(i)
}

// Int32sToUint8s converts int32 slice to uint8 slice.
func Int32sToUint8s(i []int32) ([]uint8, error) {
	var err error
	r := make([]uint8, len(i))
	for k, v := range i {
		r[k], err = Int32ToUint8(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint8s converts int32 slice to uint8 slice.
func (i Int32Slice) Uint8s() ([]uint8, error) {
	return Int32sToUint8s(i)
}

// Int32sToUint16s converts int32 slice to uint16 slice.
func Int32sToUint16s(i []int32) ([]uint16, error) {
	var err error
	r := make([]uint16, len(i))
	for k, v := range i {
		r[k], err = Int32ToUint16(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint16s converts int32 slice to uint16 slice.
func (i Int32Slice) Uint16s() ([]uint16, error) {
	return Int32sToUint16s(i)
}

// Int32sToUint32s converts int32 slice to uint32 slice.
func Int32sToUint32s(i []int32) ([]uint32, error) {
	var err error
	r := make([]uint32, len(i))
	for k, v := range i {
		r[k], err = Int32ToUint32(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint32s converts int32 slice to uint32 slice.
func (i Int32Slice) Uint32s() ([]uint32, error) {
	return Int32sToUint32s(i)
}

// Int32sToUint64s converts int32 slice to uint64 slice.
func Int32sToUint64s(i []int32) ([]uint64, error) {
	var err error
	r := make([]uint64, len(i))
	for k, v := range i {
		r[k], err = Int32ToUint64(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint64s converts int32 slice to uint64 slice.
func (i Int32Slice) Uint64s() ([]uint64, error) {
	return Int32sToUint64s(i)
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (i Int32Slice) Concat(a ...[]int32) []int32 {
	totalLen := len(i)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]int32, totalLen)
	n := copy(ret, i)
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
func (i Int32Slice) CopyWithin(target, start int, end ...int) {
	target = fixIndex(len(i), target, true)
	if target == len(i) {
		return
	}
	sub := i.Slice(start, end...)
	for k, v := range sub {
		i[target+k] = v
	}
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func (i Int32Slice) Every(fn func(curr Int32Slice, k int, v int32) bool) bool {
	for k, v := range i {
		if !fn(i, k, v) {
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
func (i Int32Slice) Fill(value int32, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(i), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		i[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (i Int32Slice) Filter(fn func(curr Int32Slice, k int, v int32) bool) []int32 {
	ret := make([]int32, 0)
	for k, v := range i {
		if fn(i, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func (i Int32Slice) Find(fn func(curr Int32Slice, k int, v int32) bool) (k int, v int32) {
	for k, v := range i {
		if fn(i, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// Includes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i Int32Slice) Includes(valueToFind int32, fromIndex ...int) bool {
	return i.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i Int32Slice) IndexOf(searchElement int32, fromIndex ...int) int {
	idx := getFromIndex(len(i), fromIndex...)
	for k, v := range i[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i Int32Slice) LastIndexOf(searchElement int32, fromIndex ...int) int {
	idx := getFromIndex(len(i), fromIndex...)
	for k := len(i) - 1; k >= idx; k-- {
		if searchElement == i[k] {
			return k
		}
	}
	return -1
}

// Map creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func (i Int32Slice) Map(fn func(curr Int32Slice, k int, v int32) int32) []int32 {
	ret := make([]int32, len(i))
	for k, v := range i {
		ret[k] = fn(i, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (i *Int32Slice) Pop() (int32, bool) {
	a := *i
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*i = a[:len(a):len(a)]
	return last, true
}

// Push adds one or more elements to the end of an slice and returns the new length of the slice.
func (i *Int32Slice) Push(element ...int32) int {
	*i = append(*i, element...)
	return len(*i)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (i *Int32Slice) PushOnce(element ...int32) int {
	a := *i
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*i = a
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
func (i Int32Slice) Reduce(
	fn func(curr Int32Slice, k int, v, accumulator int32) int32, initialValue ...int32,
) int32 {
	if len(i) == 0 {
		return 0
	}
	start := 0
	acc := i[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(i); k++ {
		acc = fn(i, k, i[k], acc)
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
func (i Int32Slice) ReduceRight(
	fn func(curr Int32Slice, k int, v, accumulator int32) int32, initialValue ...int32,
) int32 {
	if len(i) == 0 {
		return 0
	}
	end := len(i) - 1
	acc := i[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(i, k, i[k], acc)
	}
	return acc
}

// Reverse reverses an slice in place.
func (i Int32Slice) Reverse() {
	first := 0
	last := len(i) - 1
	for first < last {
		i[first], i[last] = i[last], i[first]
		first++
		last--
	}
}

// Shift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func (i *Int32Slice) Shift() (int32, bool) {
	a := *i
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*i = a[:len(a):len(a)]
	return first, true
}

// Slice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func (i Int32Slice) Slice(begin int, end ...int) []int32 {
	fixedStart, fixedEnd, ok := fixRange(len(i), begin, end...)
	if !ok {
		return []int32{}
	}
	return i[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (i Int32Slice) Some(fn func(curr Int32Slice, k int, v int32) bool) bool {
	for k, v := range i {
		if fn(i, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (i Int32Slice) Len() int {
	return len(i)
}

// Less reports whether the element with
// index m should sort before the element with index n.
func (i Int32Slice) Less(m, n int) bool {
	return i[m] < i[n]
}

// Swap swaps the elements with indexes m and n.
func (i Int32Slice) Swap(m, n int) {
	i[m], i[n] = i[n], i[m]
}

// Sort sorts the elements of an slice in place and returns the sorted slice.
func (i Int32Slice) Sort() {
	sort.Sort(i)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (i *Int32Slice) Splice(start, deleteCount int, items ...int32) {
	a := *i
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := fixRange(len(a), start, start+1+deleteCount)
	deleteCount = end - start - 1
	for k := 0; k < len(items); k++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[k]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := a[start:].Copy()
			items = items[k:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*i = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*i = a[:len(a):len(a)]
}

// Unshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func (i *Int32Slice) Unshift(element ...int32) int {
	*i = append(element, *i...)
	return len(*i)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (i *Int32Slice) UnshiftOnce(element ...int32) int {
	a := *i
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[int32]bool, len(element))
	r := make([]int32, 0, len(a)+len(element))
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
	*i = r[:len(r):len(r)]
	return len(r)
}

// Distinct creates an new slice in place set that removes the same elements
// and returns the new length of the slice.
func (i *Int32Slice) Distinct() int {
	a := (*i)[:0]
	m := make(map[int32]bool, len(a))
	for _, v := range *i {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*i = a[:n:n]
	return n
}

// RemoveOne removes the first matched elements from the slice,
// and returns the new length of the slice.
func (i *Int32Slice) RemoveOne(element ...int32) int {
	a := *i
	m := make(map[int32]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
				break
			}
		}
	}
	n := len(a)
	*i = a[:n:n]
	return n
}

// RemoveEvery removes all the elements from the slice,
// and returns the new length of the slice.
func (i *Int32Slice) RemoveEvery(element ...int32) int {
	a := *i
	m := make(map[int32]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
			}
		}
	}
	n := len(a)
	*i = a[:n:n]
	return n
}
