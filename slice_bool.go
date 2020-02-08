package ameda

// BoolSlice bool slice object
type BoolSlice []bool

// NewBoolSlice creates an BoolSlice object.
func NewBoolSlice(a []bool) *BoolSlice {
	i := BoolSlice(a)
	return &i
}

// BoolsCopy creates a copy of the bool slice.
func BoolsCopy(b []bool) []bool {
	r := make([]bool, len(b))
	copy(r, b)
	return r
}

// Copy creates a copy of the bool slice.
func (b BoolSlice) Copy() []bool {
	return BoolsCopy(b)
}

// BoolsToInterfaces converts int8 slice to interface slice.
func BoolsToInterfaces(b []bool) []interface{} {
	r := make([]interface{}, len(b))
	for k, v := range b {
		r[k] = v
	}
	return r
}

// Interfaces converts int8 slice to interface slice.
func (b BoolSlice) Interfaces() []interface{} {
	return BoolsToInterfaces(b)
}

// BoolsToStrings converts int8 slice to string slice.
func BoolsToStrings(b []bool) []string {
	r := make([]string, len(b))
	for k, v := range b {
		r[k] = BoolToString(v)
	}
	return r
}

// Strings converts int8 slice to string slice.
func (b BoolSlice) Strings() []string {
	return BoolsToStrings(b)
}

// Bools converts int8 slice to bool slice.
func (b BoolSlice) Bools() []bool {
	return []bool(b)
}

// BoolsToFloat32s converts int8 slice to float32 slice.
func BoolsToFloat32s(b []bool) []float32 {
	r := make([]float32, len(b))
	for k, v := range b {
		r[k] = BoolToFloat32(v)
	}
	return r
}

// Float32s converts int8 slice to float32 slice.
func (b BoolSlice) Float32s() []float32 {
	return BoolsToFloat32s(b)
}

// BoolsToFloat64s converts int8 slice to float64 slice.
func BoolsToFloat64s(b []bool) []float64 {
	r := make([]float64, len(b))
	for k, v := range b {
		r[k] = BoolToFloat64(v)
	}
	return r
}

// Float64s converts int8 slice to float64 slice.
func (b BoolSlice) Float64s() []float64 {
	return BoolsToFloat64s(b)
}

// BoolsToInts converts int8 slice to int slice.
func BoolsToInts(b []bool) []int {
	r := make([]int, len(b))
	for k, v := range b {
		r[k] = BoolToInt(v)
	}
	return r
}

// Ints converts int8 slice to int slice.
func (b BoolSlice) Ints() []int {
	return BoolsToInts(b)
}

// Int8s converts to []int8.
func (b BoolSlice) Int8s() []int8 {
	r := make([]int8, len(b))
	for k, v := range b {
		r[k] = BoolToInt8(v)
	}
	return r
}

// BoolsToInt16s converts int8 slice to int16 slice.
func BoolsToInt16s(b []bool) []int16 {
	r := make([]int16, len(b))
	for k, v := range b {
		r[k] = BoolToInt16(v)
	}
	return r
}

// Int16s converts int8 slice to int16 slice.
func (b BoolSlice) Int16s() []int16 {
	return BoolsToInt16s(b)
}

// BoolsToInt32s converts int8 slice to int32 slice.
func BoolsToInt32s(b []bool) []int32 {
	r := make([]int32, len(b))
	for k, v := range b {
		r[k] = BoolToInt32(v)
	}
	return r
}

// Int32s converts int8 slice to int32 slice.
func (b BoolSlice) Int32s() []int32 {
	return BoolsToInt32s(b)
}

// BoolsToInt64s converts int8 slice to int64 slice.
func BoolsToInt64s(b []bool) []int64 {
	r := make([]int64, len(b))
	for k, v := range b {
		r[k] = BoolToInt64(v)
	}
	return r
}

// Int64s converts int8 slice to int64 slice.
func (b BoolSlice) Int64s() []int64 {
	return BoolsToInt64s(b)
}

// BoolsToUints converts bool slice to uint slice.
func BoolsToUints(b []bool) []uint {
	r := make([]uint, len(b))
	for k, v := range b {
		r[k] = BoolToUint(v)
	}
	return r
}

// Uints converts bool slice to uint slice.
func (b BoolSlice) Uints() []uint {
	return BoolsToUints(b)
}

// Uint8s converts to []uint8.
func (b BoolSlice) Uint8s() []uint8 {
	r := make([]uint8, len(b))
	for k, v := range b {
		r[k] = BoolToUint8(v)
	}
	return r
}

// BoolsToUint16s converts bool slice to uint16 slice.
func BoolsToUint16s(b []bool) []uint16 {
	r := make([]uint16, len(b))
	for k, v := range b {
		r[k] = BoolToUint16(v)
	}
	return r
}

// Uint16s converts bool slice to uint16 slice.
func (b BoolSlice) Uint16s() []uint16 {
	return BoolsToUint16s(b)
}

// BoolsToUint32s converts bool slice to uint32 slice.
func BoolsToUint32s(b []bool) []uint32 {
	r := make([]uint32, len(b))
	for k, v := range b {
		r[k] = BoolToUint32(v)
	}
	return r
}

// Uint32s converts bool slice to uint32 slice.
func (b BoolSlice) Uint32s() []uint32 {
	return BoolsToUint32s(b)
}

// BoolsToUint64s converts bool slice to uint64 slice.
func BoolsToUint64s(b []bool) []uint64 {
	r := make([]uint64, len(b))
	for k, v := range b {
		r[k] = BoolToUint64(v)
	}
	return r
}

// Uint64s converts bool slice to uint64 slice.
func (b BoolSlice) Uint64s() []uint64 {
	return BoolsToUint64s(b)
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (b BoolSlice) Concat(a ...[]bool) []bool {
	totalLen := len(b)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]bool, totalLen)
	n := copy(ret, b)
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
func (b BoolSlice) CopyWithin(target, start int, end ...int) {
	target = fixIndex(len(b), target, true)
	if target == len(b) {
		return
	}
	sub := b.Slice(start, end...)
	for k, v := range sub {
		b[target+k] = v
	}
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func (b BoolSlice) Every(fn func(curr BoolSlice, k int, v bool) bool) bool {
	for k, v := range b {
		if !fn(b, k, v) {
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
func (b BoolSlice) Fill(value bool, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(b), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		b[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (b BoolSlice) Filter(fn func(curr BoolSlice, k int, v bool) bool) []bool {
	ret := make([]bool, 0)
	for k, v := range b {
		if fn(b, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func (b BoolSlice) Find(fn func(curr BoolSlice, k int, v bool) bool) (k int, v bool) {
	for k, v := range b {
		if fn(b, k, v) {
			return k, v
		}
	}
	return -1, false
}

// Includes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (b BoolSlice) Includes(valueToFind bool, fromIndex ...int) bool {
	return b.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (b BoolSlice) IndexOf(searchElement bool, fromIndex ...int) int {
	idx := getFromIndex(len(b), fromIndex...)
	for k, v := range b[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (b BoolSlice) LastIndexOf(searchElement bool, fromIndex ...int) int {
	idx := getFromIndex(len(b), fromIndex...)
	for k := len(b) - 1; k >= idx; k-- {
		if searchElement == b[k] {
			return k
		}
	}
	return -1
}

// Map creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func (b BoolSlice) Map(fn func(curr BoolSlice, k int, v bool) bool) []bool {
	ret := make([]bool, len(b))
	for k, v := range b {
		ret[k] = fn(b, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (b *BoolSlice) Pop() (elem bool, found bool) {
	a := *b
	if len(a) == 0 {
		return false, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*b = a[:len(a):len(a)]
	return last, true
}

// Push adds one or more elements to the end of an slice and returns the new length of the slice.
func (b *BoolSlice) Push(element ...bool) int {
	*b = append(*b, element...)
	return len(*b)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (b *BoolSlice) PushOnce(element ...bool) int {
	a := *b
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*b = a
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
func (b BoolSlice) Reduce(
	fn func(curr BoolSlice, k int, v bool, accumulator bool) bool, initialValue ...bool,
) bool {
	if len(b) == 0 {
		return false
	}
	start := 0
	acc := b[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(b); k++ {
		acc = fn(b, k, b[k], acc)
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
func (b BoolSlice) ReduceRight(
	fn func(curr BoolSlice, k int, v bool, accumulator bool) bool, initialValue ...bool,
) bool {
	if len(b) == 0 {
		return false
	}
	end := len(b) - 1
	acc := b[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(b, k, b[k], acc)
	}
	return acc
}

// Reverse reverses an slice in place.
func (b BoolSlice) Reverse() {
	first := 0
	last := len(b) - 1
	for first < last {
		b[first], b[last] = b[last], b[first]
		first++
		last--
	}
}

// Shift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func (b *BoolSlice) Shift() (element bool, found bool) {
	a := *b
	if len(a) == 0 {
		return false, false
	}
	first := a[0]
	a = a[1:]
	*b = a[:len(a):len(a)]
	return first, true
}

// Slice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func (b BoolSlice) Slice(begin int, end ...int) []bool {
	fixedStart, fixedEnd, ok := fixRange(len(b), begin, end...)
	if !ok {
		return []bool{}
	}
	return b[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (b BoolSlice) Some(fn func(curr BoolSlice, k int, v bool) bool) bool {
	for k, v := range b {
		if fn(b, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (b BoolSlice) Len() int {
	return len(b)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (b *BoolSlice) Splice(start, deleteCount int, items ...bool) {
	a := *b
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
			*b = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*b = a[:len(a):len(a)]
}

// Unshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func (b *BoolSlice) Unshift(element ...bool) int {
	*b = append(element, *b...)
	return len(*b)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (b *BoolSlice) UnshiftOnce(element ...bool) int {
	a := *b
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[bool]bool, len(element))
	r := make([]bool, 0, len(a)+len(element))
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
	*b = r[:len(r):len(r)]
	return len(r)
}

// Distinct creates an new slice in place set that removes the same elements
// and returns the new length of the slice.
func (b *BoolSlice) Distinct() int {
	a := (*b)[:0]
	m := make(map[bool]bool, len(a))
	for _, v := range *b {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*b = a[:n:n]
	return n
}

// RemoveOne removes the first matched elements from the slice,
// and returns the new length of the slice.
func (b *BoolSlice) RemoveOne(element ...bool) int {
	a := *b
	m := make(map[bool]bool, len(element))
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
	*b = a[:n:n]
	return n
}

// RemoveAll removes all the elements from the slice,
// and returns the new length of the slice.
func (b *BoolSlice) RemoveAll(element ...bool) int {
	a := *b
	m := make(map[bool]bool, len(element))
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
	*b = a[:n:n]
	return n
}
