package iter

type Chain[T comparable] struct {
	other Iterator[T]
}

func newChain[T comparable](other Iterator[T]) *Chain[T] {
	return &Chain[T]{other: other}
}
