package iter

type Fuse[T comparable] struct {
	iter Iterator[T]
}

func newFuse[T comparable](iter Iterator[T]) *Fuse[T] {
	return &Fuse[T]{iter: iter}
}
