package iter

type Chain[T any] struct {
	other Iterator[T]
}

func newChain[T any](other Iterator[T]) *Chain[T] {
	return &Chain[T]{other: other}
}
