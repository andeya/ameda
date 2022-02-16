package iter

type Map[T any] struct {
	iter Iterator[T]
	f    func(T) any
}

func newMap[T any](iter Iterator[T], f func(T) any) *Map[T] {
	return &Map[T]{iter: iter, f: f}
}
