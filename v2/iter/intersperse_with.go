package iter

type IntersperseWith[T comparable] struct {
	iter      Iterator[T]
	separator func() T
}

func newIntersperseWith[T comparable](iter Iterator[T], separator func() T) *IntersperseWith[T] {
	return &IntersperseWith[T]{iter: iter, separator: separator}
}
