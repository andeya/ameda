package iter

type Intersperse[T comparable] struct {
	iter      Iterator[T]
	separator T
}

func newIntersperse[T comparable](iter Iterator[T], separator T) *Intersperse[T] {
	return &Intersperse[T]{iter: iter, separator: separator}
}

