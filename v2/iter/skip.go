package iter

type Skip[T any] struct {
	iter Iterator[T]
	n    int
}

func newSkip[T any](iter Iterator[T], n int) *Skip[T] {
	return &Skip[T]{iter: iter, n: n}
}
