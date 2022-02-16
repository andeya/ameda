package iter

type FlatMap[T any] struct {
	iter Iterator[T]
	f    func(T) IntoIterator[any]
}

func newFlatMap[T any](iter Iterator[T], f func(T) IntoIterator[any]) *FlatMap[T] {
	return &FlatMap[T]{iter: iter, f: f}
}
