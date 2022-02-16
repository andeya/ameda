package iter

type IntersperseWith[T any] struct {
	iter      Iterator[T]
	separator func() T
}

func newIntersperseWith[T any](iter Iterator[T], separator func() T) *IntersperseWith[T] {
	return &IntersperseWith[T]{iter: iter, separator: separator}
}
