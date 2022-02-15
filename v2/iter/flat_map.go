package iter

type FlatMap[T comparable] struct {
	iter Iterator[T]
	f    func(T) IntoIterator[any]
}

func newFlatMap[T comparable](iter Iterator[T], f func(T) IntoIterator[any]) *FlatMap[T] {
	return &FlatMap[T]{iter: iter, f: f}
}
