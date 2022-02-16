package iter

type Take[T any] struct {
	iter Iterator[T]
	n    int
}

func newTake[T any](iter Iterator[T], n int) *Take[T] {
	return &Take[T]{iter: iter, n: n}
}
