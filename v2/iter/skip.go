package iter

type Skip[T comparable] struct {
	iter Iterator[T]
	n    int
}

func newSkip[T comparable](iter Iterator[T], n int) *Skip[T] {
	return &Skip[T]{iter: iter, n: n}
}
