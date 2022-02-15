package iter

type Enumerate[T comparable] struct {
	iter Iterator[T]
}

func newEnumerate[T comparable](iter Iterator[T]) *Enumerate[T] {
	return &Enumerate[T]{iter: iter}
}
