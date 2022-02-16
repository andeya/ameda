package iter

type Enumerate[T any] struct {
	iter Iterator[T]
}

func newEnumerate[T any](iter Iterator[T]) *Enumerate[T] {
	return &Enumerate[T]{iter: iter}
}
