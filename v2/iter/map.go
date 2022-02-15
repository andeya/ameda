package iter

type Map[T comparable] struct {
	iter Iterator[T]
	f    func(T) any
}

func newMap[T comparable](iter Iterator[T], f func(T) any) *Map[T] {
	return &Map[T]{iter: iter, f: f}
}
