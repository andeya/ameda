package iter

type SkipWhile[T any] struct {
	iter Iterator[T]
	f    func(T) bool
}

func newSkipWhile[T any](iter Iterator[T], f func(T) bool) *SkipWhile[T] {
	return &SkipWhile[T]{iter: iter, f: f}
}
