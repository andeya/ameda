package iter

type Flatten[T any, B any] struct {
	iter Iterator[T]
}

func newFlatten[T any, B any](iter Iterator[T]) *Flatten[T, B] {
	return &Flatten[T, B]{iter: iter}
}
