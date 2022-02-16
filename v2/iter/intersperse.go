package iter

type Intersperse[T any] struct {
	iter      Iterator[T]
	separator T
}

func newIntersperse[T any](iter Iterator[T], separator T) *Intersperse[T] {
	return &Intersperse[T]{iter: iter, separator: separator}
}

