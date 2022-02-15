package iter

type Zip[T comparable] struct {
	other Iterator[T]
}

func newZip[T comparable](other Iterator[T]) *Zip[T] {
	return &Zip[T]{other: other}
}
