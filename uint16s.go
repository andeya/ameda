package ameda

// OneUint16 try to return the first element, otherwise return zero value.
func OneUint16(u []uint16) uint16 {
	if len(u) > 0 {
		return u[0]
	}
	return 0
}

// Uint16sCopy creates a copy of the uint16 slice.
func Uint16sCopy(u []uint16) []uint16 {
	b := make([]uint16, len(u))
	copy(b, u)
	return b
}

// Uint16sToInterfaces converts uint16 slice to interface slice.
func Uint16sToInterfaces(u []uint16) []interface{} {
	r := make([]interface{}, len(u))
	for k, v := range u {
		r[k] = Uint16ToInterface(v)
	}
	return r
}

// Uint16sToStrings converts uint16 slice to string slice.
func Uint16sToStrings(u []uint16) []string {
	r := make([]string, len(u))
	for k, v := range u {
		r[k] = Uint16ToString(v)
	}
	return r
}

// Uint16sToBools converts uint16 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func Uint16sToBools(u []uint16) []bool {
	r := make([]bool, len(u))
	for k, v := range u {
		r[k] = Uint16ToBool(v)
	}
	return r
}

// Uint16sToFloat32s converts uint16 slice to float32 slice.
func Uint16sToFloat32s(u []uint16) []float32 {
	r := make([]float32, len(u))
	for k, v := range u {
		r[k] = Uint16ToFloat32(v)
	}
	return r
}

// Uint16sToFloat64s converts uint16 slice to float64 slice.
func Uint16sToFloat64s(u []uint16) []float64 {
	r := make([]float64, len(u))
	for k, v := range u {
		r[k] = Uint16ToFloat64(v)
	}
	return r
}

// Uint16sToInts converts uint16 slice to int slice.
func Uint16sToInts(u []uint16) []int {
	r := make([]int, len(u))
	for k, v := range u {
		r[k] = Uint16ToInt(v)
	}
	return r
}

// Uint16sToInt8s converts uint16 slice to int8 slice.
func Uint16sToInt8s(u []uint16) ([]int8, error) {
	var err error
	r := make([]int8, len(u))
	for k, v := range u {
		r[k], err = Uint16ToInt8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint16sToInt16s converts uint16 slice to int16 slice.
func Uint16sToInt16s(u []uint16) ([]int16, error) {
	var err error
	r := make([]int16, len(u))
	for k, v := range u {
		r[k], err = Uint16ToInt16(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint16sToInt32s converts uint16 slice to int32 slice.
func Uint16sToInt32s(u []uint16) []int32 {
	r := make([]int32, len(u))
	for k, v := range u {
		r[k] = Uint16ToInt32(v)
	}
	return r
}

// Uint16sToInt64s converts uint16 slice to int64 slice.
func Uint16sToInt64s(u []uint16) []int64 {
	r := make([]int64, len(u))
	for k, v := range u {
		r[k] = Uint16ToInt64(v)
	}
	return r
}

// Uint16sToUints converts uint16 slice to uint slice.
func Uint16sToUints(u []uint16) []uint {
	r := make([]uint, len(u))
	for k, v := range u {
		r[k] = Uint16ToUint(v)
	}
	return r
}

// Uint16sToUint8s converts uint16 slice to uint8 slice.
func Uint16sToUint8s(u []uint16) ([]uint8, error) {
	var err error
	r := make([]uint8, len(u))
	for k, v := range u {
		r[k], err = Uint16ToUint8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint16sToUint32s converts uint16 slice to uint32 slice.
func Uint16sToUint32s(u []uint16) []uint32 {
	r := make([]uint32, len(u))
	for k, v := range u {
		r[k] = Uint16ToUint32(v)
	}
	return r
}

// Uint16sToUint64s converts uint16 slice to uint64 slice.
func Uint16sToUint64s(u []uint16) []uint64 {
	r := make([]uint64, len(u))
	for k, v := range u {
		r[k] = Uint16ToUint64(v)
	}
	return r
}

// Uint16sConcat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func Uint16sConcat(u ...[]uint16) []uint16 {
	var totalLen int
	for _, v := range u {
		totalLen += len(v)
	}
	ret := make([]uint16, totalLen)
	dst := ret
	for _, v := range u {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// Uint16sCopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func Uint16sCopyWithin(u []uint16, target, start int, end ...int) {
	target = fixIndex(len(u), target, true)
	if target == len(u) {
		return
	}
	sub := Uint16sSlice(u, start, end...)
	for k, v := range sub {
		u[target+k] = v
	}
}

// Uint16sEvery tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func Uint16sEvery(u []uint16, fn func(u []uint16, k int, v uint16) bool) bool {
	for k, v := range u {
		if !fn(u, k, v) {
			return false
		}
	}
	return true
}

// Uint16sFill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func Uint16sFill(u []uint16, value uint16, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(u), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		u[k] = value
	}
}

// Uint16sFilter creates a new slice with all elements that pass the test implemented by the provided function.
func Uint16sFilter(u []uint16, fn func(u []uint16, k int, v uint16) bool) []uint16 {
	ret := make([]uint16, 0)
	for k, v := range u {
		if fn(u, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Uint16sFind returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func Uint16sFind(u []uint16, fn func(u []uint16, k int, v uint16) bool) (k int, v uint16) {
	for k, v := range u {
		if fn(u, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// Uint16sIncludes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func Uint16sIncludes(u []uint16, valueToFind uint16, fromIndex ...int) bool {
	return Uint16sIndexOf(u, valueToFind, fromIndex...) > -1
}

// Uint16sIndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func Uint16sIndexOf(u []uint16, searchElement uint16, fromIndex ...int) int {
	idx := getFromIndex(len(u), fromIndex...)
	for k, v := range u[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// Uint16sLastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func Uint16sLastIndexOf(u []uint16, searchElement uint16, fromIndex ...int) int {
	idx := getFromIndex(len(u), fromIndex...)
	for k := len(u) - 1; k >= idx; k-- {
		if searchElement == u[k] {
			return k
		}
	}
	return -1
}

// Uint16sMap creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func Uint16sMap(u []uint16, fn func(u []uint16, k int, v uint16) uint16) []uint16 {
	ret := make([]uint16, len(u))
	for k, v := range u {
		ret[k] = fn(u, k, v)
	}
	return ret
}

// Uint16sPop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func Uint16sPop(u *[]uint16) (uint16, bool) {
	a := *u
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*u = a[:len(a):len(a)]
	return last, true
}

// Uint16sPush adds one or more elements to the end of an slice and returns the new length of the slice.
func Uint16sPush(u *[]uint16, element ...uint16) int {
	*u = append(*u, element...)
	return len(*u)
}

// Uint16sPushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func Uint16sPushOnce(u *[]uint16, element ...uint16) int {
	a := *u
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*u = a
	return len(a)
}

// Uint16sReduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func Uint16sReduce(
	u []uint16,
	fn func(u []uint16, k int, v, accumulator uint16) uint16, initialValue ...uint16,
) uint16 {
	if len(u) == 0 {
		return 0
	}
	start := 0
	acc := u[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(u); k++ {
		acc = fn(u, k, u[k], acc)
	}
	return acc
}

// Uint16sReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func Uint16sReduceRight(
	u []uint16,
	fn func(u []uint16, k int, v, accumulator uint16) uint16, initialValue ...uint16,
) uint16 {
	if len(u) == 0 {
		return 0
	}
	end := len(u) - 1
	acc := u[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(u, k, u[k], acc)
	}
	return acc
}

// Uint16sReverse reverses an slice in place.
func Uint16sReverse(u []uint16) {
	first := 0
	last := len(u) - 1
	for first < last {
		u[first], u[last] = u[last], u[first]
		first++
		last--
	}
}

// Uint16sShift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func Uint16sShift(u *[]uint16) (uint16, bool) {
	a := *u
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*u = a[:len(a):len(a)]
	return first, true
}

// Uint16sSlice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func Uint16sSlice(u []uint16, begin int, end ...int) []uint16 {
	fixedStart, fixedEnd, ok := fixRange(len(u), begin, end...)
	if !ok {
		return []uint16{}
	}
	return Uint16sCopy(u[fixedStart:fixedEnd])
}

// Uint16sSome tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func Uint16sSome(u []uint16, fn func(u []uint16, k int, v uint16) bool) bool {
	for k, v := range u {
		if fn(u, k, v) {
			return true
		}
	}
	return false
}

// Uint16sSplice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func Uint16sSplice(u *[]uint16, start, deleteCount int, items ...uint16) {
	a := *u
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
			lastSlice := Uint16sCopy(a[start:])
			items = items[k:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*u = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*u = a[:len(a):len(a)]
}

// Uint16sUnshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func Uint16sUnshift(u *[]uint16, element ...uint16) int {
	*u = append(element, *u...)
	return len(*u)
}

// Uint16sUnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func Uint16sUnshiftOnce(u *[]uint16, element ...uint16) int {
	a := *u
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[uint16]bool, len(element))
	r := make([]uint16, 0, len(a)+len(element))
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
	*u = r[:len(r):len(r)]
	return len(r)
}

// Uint16sDistinct creates a new slice in place set that removes the same elements
// and returns the new length of the slice.
func Uint16sDistinct(u *[]uint16) int {
	a := (*u)[:0]
	m := make(map[uint16]bool, len(a))
	for _, v := range *u {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*u = a[:n:n]
	return n
}

// Uint16sRemoveFirst removes the first matched elements from the slice,
// and returns the new length of the slice.
func Uint16sRemoveFirst(u *[]uint16, element ...uint16) int {
	a := *u
	m := make(map[uint16]bool, len(element))
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
	*u = a[:n:n]
	return n
}

// Uint16sRemoveEvery removes all the elements from the slice,
// and returns the new length of the slice.
func Uint16sRemoveEvery(u *[]uint16, element ...uint16) int {
	a := *u
	m := make(map[uint16]bool, len(element))
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
	*u = a[:n:n]
	return n
}
