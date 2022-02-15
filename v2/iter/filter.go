package iter

type Filter[T comparable] struct {
	iter Iterator[T]
	f    func(T) bool
}

func newFilter[T comparable](iter Iterator[T], f func(T) bool) *Filter[T] {
	return &Filter[T]{iter: iter, f: f}
}
