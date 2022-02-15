package iter

type Flatten[T comparable, B comparable] struct {
	iter Iterator[T]
}

func newFlatten[T comparable, B comparable](iter Iterator[T]) *Flatten[T, B] {
	return &Flatten[T, B]{iter: iter}
}
