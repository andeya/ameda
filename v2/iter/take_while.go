package iter

type TakeWhile[T any] struct {
	iter Iterator[T]
	f    func(T) bool
}

func newTakeWhile[T any](iter Iterator[T], f func(T) bool) *TakeWhile[T] {
	return &TakeWhile[T]{iter: iter, f: f}
}
