package iter

type Fuse[T any] struct {
	iter Iterator[T]
}

func newFuse[T any](iter Iterator[T]) *Fuse[T] {
	return &Fuse[T]{iter: iter}
}
