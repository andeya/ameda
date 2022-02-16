package iter

type Zip[T any] struct {
	other Iterator[T]
}

func newZip[T any](other Iterator[T]) *Zip[T] {
	return &Zip[T]{other: other}
}
