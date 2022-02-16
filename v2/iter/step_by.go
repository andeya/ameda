package iter

type StepBy[T any] struct {
}

func newStepBy[T any](iter *baseIterator[T], step int) *StepBy[T] {
	return &StepBy[T]{}
}
