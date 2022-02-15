package iter

type Take[T comparable] struct {
	iter Iterator[T]
	n    int
}

func newTake[T comparable](iter Iterator[T], n int) *Take[T] {
	return &Take[T]{iter: iter, n: n}
}
