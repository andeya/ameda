package ameda

// InterfaceSlice interface slice object
type InterfaceSlice []interface{}

// NewInterfaceSlice creates an InterfaceSlice object.
func NewInterfaceSlice(a []interface{}) *InterfaceSlice {
	i := InterfaceSlice(a)
	return &i
}

// InterfacesCopy creates a copy of the interface slice.
func InterfacesCopy(i []interface{}) []interface{} {
	b := make([]interface{}, len(i))
	copy(b, i)
	return b
}

// Copy creates a copy of the interface slice.
func (i InterfaceSlice) Copy() []interface{} {
	return InterfacesCopy(i)
}

// Interfaces converts interface slice to interface slice.
func (i InterfaceSlice) Interfaces() []interface{} {
	return []interface{}(i)
}

// InterfacesToStrings converts interface slice to string slice.
func InterfacesToStrings(i []interface{}) []string {
	r := make([]string, len(i))
	for k, v := range i {
		r[k] = InterfaceToString(v)
	}
	return r
}

// Strings converts interface slice to string slice.
func (i InterfaceSlice) Strings() []string {
	return InterfacesToStrings(i)
}

// InterfacesToBools converts interface slice to bool slice.
// NOTE:
//  0 is false, other numbers are true
func InterfacesToBools(i []interface{}) ([]bool, error) {
	var err error
	r := make([]bool, len(i))
	for k, v := range i {
		r[k], err = InterfaceToBool(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Bools converts interface slice to bool slice.
// NOTE:
//  0 is false, other numbers are true
func (i InterfaceSlice) Bools() ([]bool, error) {
	return InterfacesToBools(i)
}

// InterfacesToFloat32s converts interface slice to float32 slice.
func InterfacesToFloat32s(i []interface{}) ([]float32, error) {
	var err error
	r := make([]float32, len(i))
	for k, v := range i {
		r[k], err = InterfaceToFloat32(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Float32s converts interface slice to float32 slice.
func (i InterfaceSlice) Float32s() ([]float32, error) {
	return InterfacesToFloat32s(i)
}

// InterfacesToFloat64s converts interface slice to float64 slice.
func InterfacesToFloat64s(i []interface{}) ([]float64, error) {
	var err error
	r := make([]float64, len(i))
	for k, v := range i {
		r[k], err = InterfaceToFloat64(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Float64s converts interface slice to float64 slice.
func (i InterfaceSlice) Float64s() ([]float64, error) {
	return InterfacesToFloat64s(i)
}

// InterfacesToInts converts interface slice to int slice.
func InterfacesToInts(i []interface{}) ([]int, error) {
	var err error
	r := make([]int, len(i))
	for k, v := range i {
		r[k], err = InterfaceToInt(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Ints converts interface slice to int slice.
func (i InterfaceSlice) Ints() ([]int, error) {
	return InterfacesToInts(i)
}

// InterfacesToInt8s converts interface slice to int8 slice.
func InterfacesToInt8s(i []interface{}) ([]int8, error) {
	var err error
	r := make([]int8, len(i))
	for k, v := range i {
		r[k], err = InterfaceToInt8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8s converts interface slice to int8 slice.
func (i InterfaceSlice) Int8s() ([]int8, error) {
	return InterfacesToInt8s(i)
}

// InterfacesToInt16s converts interface slice to int16 slice.
func InterfacesToInt16s(i []interface{}) ([]int16, error) {
	var err error
	r := make([]int16, len(i))
	for k, v := range i {
		r[k], err = InterfaceToInt16(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int16s converts interface slice to int16 slice.
func (i InterfaceSlice) Int16s() ([]int16, error) {
	return InterfacesToInt16s(i)
}

// InterfacesToInt32s converts interface slice to int32 slice.
func InterfacesToInt32s(i []interface{}) ([]int32, error) {
	var err error
	r := make([]int32, len(i))
	for k, v := range i {
		r[k], err = InterfaceToInt32(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int32s converts interface slice to int32 slice.
func (i InterfaceSlice) Int32s() ([]int32, error) {
	return InterfacesToInt32s(i)
}

// InterfacesToInt64s converts interface slice to int64 slice.
func InterfacesToInt64s(i []interface{}) ([]int64, error) {
	var err error
	r := make([]int64, len(i))
	for k, v := range i {
		r[k], err = InterfaceToInt64(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int64s converts to []int64.
func (i InterfaceSlice) Int64s() ([]int64, error) {
	return InterfacesToInt64s(i)
}

// InterfacesToUints converts interface slice to uint slice.
func InterfacesToUints(i []interface{}) ([]uint, error) {
	var err error
	r := make([]uint, len(i))
	for k, v := range i {
		r[k], err = InterfaceToUint(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uints converts interface slice to uint slice.
func (i InterfaceSlice) Uints() ([]uint, error) {
	return InterfacesToUints(i)
}

// InterfacesToUint8s converts interface slice to uint8 slice.
func InterfacesToUint8s(i []interface{}) ([]uint8, error) {
	var err error
	r := make([]uint8, len(i))
	for k, v := range i {
		r[k], err = InterfaceToUint8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint8s converts interface slice to uint8 slice.
func (i InterfaceSlice) Uint8s() ([]uint8, error) {
	return InterfacesToUint8s(i)
}

// InterfacesToUint16s converts interface slice to uint16 slice.
func InterfacesToUint16s(i []interface{}) ([]uint16, error) {
	var err error
	r := make([]uint16, len(i))
	for k, v := range i {
		r[k], err = InterfaceToUint16(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint16s converts interface slice to uint16 slice.
func (i InterfaceSlice) Uint16s() ([]uint16, error) {
	return InterfacesToUint16s(i)
}

// InterfacesToUint32s converts interface slice to uint32 slice.
func InterfacesToUint32s(i []interface{}) ([]uint32, error) {
	var err error
	r := make([]uint32, len(i))
	for k, v := range i {
		r[k], err = InterfaceToUint32(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint32s converts interface slice to uint32 slice.
func (i InterfaceSlice) Uint32s() ([]uint32, error) {
	return InterfacesToUint32s(i)
}

// InterfacesToUint64s converts interface slice to uint64 slice.
func InterfacesToUint64s(i []interface{}) ([]uint64, error) {
	var err error
	r := make([]uint64, len(i))
	for k, v := range i {
		r[k], err = InterfaceToUint64(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Uint64s converts interface slice to uint64 slice.
func (i InterfaceSlice) Uint64s() ([]uint64, error) {
	return InterfacesToUint64s(i)
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (i InterfaceSlice) Concat(a ...[]interface{}) []interface{} {
	totalLen := len(i)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]interface{}, totalLen)
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
func (i InterfaceSlice) CopyWithin(target, start int, end ...int) {
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
func (i InterfaceSlice) Every(fn func(curr InterfaceSlice, k int, v interface{}) bool) bool {
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
func (i InterfaceSlice) Fill(value []interface{}, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(i), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		i[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (i InterfaceSlice) Filter(fn func(curr InterfaceSlice, k int, v interface{}) bool) []interface{} {
	ret := make([]interface{}, 0)
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
func (i InterfaceSlice) Find(fn func(curr InterfaceSlice, k int, v interface{}) bool) (k int, v interface{}) {
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
func (i InterfaceSlice) Includes(valueToFind int64, fromIndex ...int) bool {
	return i.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i InterfaceSlice) IndexOf(searchElement int64, fromIndex ...int) int {
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
func (i InterfaceSlice) LastIndexOf(searchElement int64, fromIndex ...int) int {
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
func (i InterfaceSlice) Map(fn func(curr InterfaceSlice, k int, v interface{}) int64) []int64 {
	ret := make([]int64, len(i))
	for k, v := range i {
		ret[k] = fn(i, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (i *InterfaceSlice) Pop() (interface{}, bool) {
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
func (i *InterfaceSlice) Push(element ...interface{}) int {
	*i = append(*i, element...)
	return len(*i)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (i *InterfaceSlice) PushOnce(element ...interface{}) int {
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
func (i InterfaceSlice) Reduce(
	fn func(curr InterfaceSlice, k int, v, accumulator interface{}) interface{}, initialValue ...interface{},
) interface{} {
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
func (i InterfaceSlice) ReduceRight(
	fn func(curr InterfaceSlice, k int, v, accumulator interface{}) interface{}, initialValue ...interface{},
) interface{} {
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
func (i InterfaceSlice) Reverse() {
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
func (i *InterfaceSlice) Shift() (interface{}, bool) {
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
func (i InterfaceSlice) Slice(begin int, end ...int) []interface{} {
	fixedStart, fixedEnd, ok := fixRange(len(i), begin, end...)
	if !ok {
		return []interface{}{}
	}
	return i[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (i InterfaceSlice) Some(fn func(curr InterfaceSlice, k int, v interface{}) bool) bool {
	for k, v := range i {
		if fn(i, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (i InterfaceSlice) Len() int {
	return len(i)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (i *InterfaceSlice) Splice(start, deleteCount int, items ...interface{}) {
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
func (i *InterfaceSlice) Unshift(element ...interface{}) int {
	*i = append(element, *i...)
	return len(*i)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (i *InterfaceSlice) UnshiftOnce(element ...interface{}) int {
	a := *i
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[interface{}]bool, len(element))
	r := make([]interface{}, 0, len(a)+len(element))
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
func (i *InterfaceSlice) Distinct() int {
	a := (*i)[:0]
	m := make(map[interface{}]bool, len(a))
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
func (i *InterfaceSlice) RemoveOne(element ...interface{}) int {
	a := *i
	m := make(map[interface{}]bool, len(element))
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
func (i *InterfaceSlice) RemoveEvery(element ...interface{}) int {
	a := *i
	m := make(map[interface{}]bool, len(element))
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
