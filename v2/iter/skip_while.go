package iter

type SkipWhile[T comparable] struct {
	iter Iterator[T]
	f    func(T) bool
}

func newSkipWhile[T comparable](iter Iterator[T], f func(T) bool) *SkipWhile[T] {
	return &SkipWhile[T]{iter: iter, f: f}
}
