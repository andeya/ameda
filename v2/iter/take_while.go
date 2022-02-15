package iter

type TakeWhile[T comparable] struct {
	iter Iterator[T]
	f    func(T) bool
}

func newTakeWhile[T comparable](iter Iterator[T], f func(T) bool) *TakeWhile[T] {
	return &TakeWhile[T]{iter: iter, f: f}
}
