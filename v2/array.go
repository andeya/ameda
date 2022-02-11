package ameda

// VecOne try to return the first element, otherwise return zero value.
func VecOne[T any](s []T) T {
	if len(s) > 0 {
		return s[0]
	}
	return Zero[T]()
}

// VecCopy creates a copy of the slice.
func VecCopy[T any](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)
	return r
}

// VecCopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. VecCopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, VecCopyWithin will copy until the last index (default to len(s)).
func VecCopyWithin[T any](s []T, target, start int, end ...int) {
	target = fixIndex(len(s), target, true)
	if target == len(s) {
		return
	}
	sub := VecSlice(s, start, end...)
	for i, v := range sub {
		s[target+i] = v
	}
}

// VecEvery tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func VecEvery[T any](s []T, fn func(s []T, k int, v T) bool) bool {
	for k, v := range s {
		if !fn(s, k, v) {
			return false
		}
	}
	return true
}

// VecFill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. VecCopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, VecCopyWithin will copy until the last index (default to len(s)).
func VecFill[T any](s []T, value T, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(s), start, end...)
	if !ok {
		return
	}
	for i := fixedStart; i < fixedEnd; i++ {
		s[i] = value
	}
}

// VecFilter creates a new slice with all elements that pass the test implemented by the provided function.
func VecFilter[T any](s []T, fn func(s []T, k int, v T) bool) []T {
	ret := make([]T, 0)
	for k, v := range s {
		if fn(s, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// VecFind returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func VecFind[T any](s []T, fn func(s []T, k int, v T) bool) (k int, v T) {
	for k, v := range s {
		if fn(s, k, v) {
			return k, v
		}
	}
	return -1, Zero[T]()
}

// VecIncludes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func VecIncludes[T comparable](s []T, valueToFind T, fromIndex ...int) bool {
	return VecIndexOf(s, valueToFind, fromIndex...) > -1
}

// VecIndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func VecIndexOf[T comparable](s []T, searchElement T, fromIndex ...int) int {
	idx := getFromIndex(len(s), fromIndex...)
	for k, v := range s[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// VecLastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func VecLastIndexOf[T comparable](s []T, searchElement T, fromIndex ...int) int {
	idx := getFromIndex(len(s), fromIndex...)
	for i := len(s) - 1; i >= idx; i-- {
		if searchElement == s[i] {
			return i
		}
	}
	return -1
}

// VecMap creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func VecMap[T any](s []T, fn func(s []T, k int, v T) T) []T {
	ret := make([]T, len(s))
	for k, v := range s {
		ret[k] = fn(s, k, v)
	}
	return ret
}

// VecPop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func VecPop[T any](s *[]T) (T, bool) {
	a := *s
	if len(a) == 0 {
		return Zero[T](), false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*s = a[:len(a):len(a)]
	return last, true
}

// VecPush adds one or more elements to the end of an slice and returns the new length of the slice.
func VecPush[T any](s *[]T, element ...T) int {
	*s = append(*s, element...)
	return len(*s)
}

// VecPushDistinct adds one or more new elements that do not exist in the current slice at the end.
func VecPushDistinct[T comparable](s []T, element ...T) []T {
L:
	for _, v := range element {
		for _, vv := range s {
			if vv == v {
				continue L
			}
		}
		s = append(s, v)
	}
	return s
}

// VecReduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func VecReduce[T any](s []T, fn func(s []T, k int, v, accumulator T) T, initialValue ...T) T {
	if len(s) == 0 {
		return Zero[T]()
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

// VecReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func VecReduceRight[T any](s []T, fn func(s []T, k int, v, accumulator T) T, initialValue ...T) T {
	if len(s) == 0 {
		return Zero[T]()
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

// VecReverse reverses an slice in place.
func VecReverse[T any](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// VecShift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func VecShift[T any](s *[]T) (T, bool) {
	a := *s
	if len(a) == 0 {
		return Zero[T](), false
	}
	first := a[0]
	a = a[1:]
	*s = a[:len(a):len(a)]
	return first, true
}

// VecSlice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func VecSlice[T any](s []T, begin int, end ...int) []T {
	fixedStart, fixedEnd, ok := fixRange(len(s), begin, end...)
	if !ok {
		return []T{}
	}
	return VecCopy[T](s[fixedStart:fixedEnd])
}

// VecSome tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func VecSome[T any](s []T, fn func(s []T, k int, v T) bool) bool {
	for k, v := range s {
		if fn(s, k, v) {
			return true
		}
	}
	return false
}

// VecSplice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func VecSplice[T any](s *[]T, start, deleteCount int, items ...T) {
	a := *s
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := fixRange(len(a), start, start+1+deleteCount)
	deleteCount = end - start - 1
	for i := 0; i < len(items); i++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[i]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := VecCopy[T](a[start:])
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

// VecUnshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func VecUnshift[T any](s *[]T, element ...T) int {
	*s = append(element, *s...)
	return len(*s)
}

// VecUnshiftDistinct adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func VecUnshiftDistinct[T comparable](s *[]T, element ...T) int {
	a := *s
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[T]bool, len(element))
	r := make([]T, 0, len(a)+len(element))
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

// VecRemoveFirst removes the first matched elements from the slice,
// and returns the new length of the slice.
func VecRemoveFirst[T comparable](p *[]T, elements ...T) int {
	a := *p
	m := make(map[interface{}]struct{}, len(elements))
	for _, element := range elements {
		if _, ok := m[element]; ok {
			continue
		}
		m[element] = struct{}{}
		for k, v := range a {
			if v == element {
				a = append(a[:k], a[k+1:]...)
				break
			}
		}
	}
	n := len(a)
	*p = a[:n:n]
	return n
}

// VecRemoveEvery removes all the elements from the slice,
// and returns the new length of the slice.
func VecRemoveEvery[T comparable](p *[]T, elements ...T) int {
	a := *p
	m := make(map[interface{}]struct{}, len(elements))
	for _, element := range elements {
		if _, ok := m[element]; ok {
			continue
		}
		m[element] = struct{}{}
		for i := 0; i < len(a); i++ {
			if a[i] == element {
				a = append(a[:i], a[i+1:]...)
				i--
			}
		}
	}
	n := len(a)
	*p = a[:n:n]
	return n
}

// VecConcat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func VecConcat[T any](s ...[]T) []T {
	var totalLen int
	for _, v := range s {
		totalLen += len(v)
	}
	ret := make([]T, totalLen)
	dst := ret
	for _, v := range s {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// VecIntersect calculates intersection of two or more slices,
// and returns the count of each element.
func VecIntersect[T comparable](s ...[]T) (intersectCount map[T]int) {
	if len(s) == 0 {
		return nil
	}
	for _, v := range s {
		if len(v) == 0 {
			return nil
		}
	}
	counts := make([]map[T]int, len(s))
	for k, v := range s {
		counts[k] = vecDistinct(v, nil)
	}
	intersectCount = counts[0]
L:
	for k, v := range intersectCount {
		for _, c := range counts[1:] {
			v2 := c[k]
			if v2 == 0 {
				delete(intersectCount, k)
				continue L
			}
			if v > v2 {
				v = v2
			}
		}
		intersectCount[k] = v
	}
	return intersectCount
}

// VecDistinct calculates the count of each different element,
// and only saves these different elements in place if changeSlice is true.
func VecDistinct[T comparable](s *[]T, changeSlice bool) (distinctCount map[T]int) {
	if !changeSlice {
		return vecDistinct(*s, nil)
	}
	a := (*s)[:0]
	distinctCount = vecDistinct(*s, &a)
	n := len(distinctCount)
	*s = a[:n:n]
	return distinctCount
}

func vecDistinct[T comparable](src []T, dst *[]T) map[T]int {
	m := make(map[T]int, len(src))
	if dst == nil {
		for _, v := range src {
			n := m[v]
			m[v] = n + 1
		}
	} else {
		a := *dst
		for _, v := range src {
			n := m[v]
			m[v] = n + 1
			if n == 0 {
				a = append(a, v)
			}
		}
		*dst = a
	}
	return m
}

// SetsUnion calculates between multiple collections: set1 ∪ set2 ∪ others...
// This method does not change the existing slices, but instead returns a new slice.
func SetsUnion[T comparable](set1, set2 []T, others ...[]T) []T {
	m := make(map[T]struct{}, len(set1)+len(set2))
	r := make([]T, 0, len(m))
	for _, set := range append([][]T{set1, set2}, others...) {
		for _, v := range set {
			_, ok := m[v]
			if ok {
				continue
			}
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}

// SetsIntersect calculates between multiple collections: set1 ∩ set2 ∩ others...
// This method does not change the existing slices, but instead returns a new slice.
func SetsIntersect[T comparable](set1, set2 []T, others ...[]T) []T {
	sets := append([][]T{set2}, others...)
	setsCount := make([]map[T]int, len(sets))
	for k, v := range sets {
		setsCount[k] = vecDistinct(v, nil)
	}
	m := make(map[T]struct{}, len(set1))
	r := make([]T, 0, len(m))
L:
	for _, v := range set1 {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		for _, m2 := range setsCount {
			if m2[v] == 0 {
				continue L
			}
		}
		r = append(r, v)
	}
	return r
}

// SetsDifference calculates between multiple collections: set1 - set2 - others...
// This method does not change the existing slices, but instead returns a new slice.
func SetsDifference[T comparable](set1, set2 []T, others ...[]T) []T {
	m := make(map[T]struct{}, len(set1))
	r := make([]T, 0, len(set1))
	sets := append([][]T{set2}, others...)
	for _, v := range sets {
		inter := SetsIntersect(set1, v)
		for _, v := range inter {
			m[v] = struct{}{}
		}
	}
	for _, v := range set1 {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}
