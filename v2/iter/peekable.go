package iter

type Peekable[T any] struct {
	iter Iterator[T]
}

func newPeekable[T any](iter Iterator[T]) *Peekable[T] {
	return &Peekable[T]{iter: iter}
}
