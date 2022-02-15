package iter

type StepBy[T comparable] struct {
}

func newStepBy[T comparable](iter *baseIterator[T], step int) *StepBy[T] {
	return &StepBy[T]{}
}
