package iter

type Filter[T any] struct {
	iter Iterator[T]
	f    func(T) bool
}

func newFilter[T any](iter Iterator[T], f func(T) bool) *Filter[T] {
	return &Filter[T]{iter: iter, f: f}
}
