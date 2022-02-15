package iter

type Peekable[T comparable] struct {
	iter Iterator[T]
}

func newPeekable[T comparable](iter Iterator[T]) *Peekable[T] {
	return &Peekable[T]{iter: iter}
}
