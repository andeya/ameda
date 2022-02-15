package result

type Result[T any] struct {
	ok  T
	err error
}

func Ok[T any](ok T) Result[T] {
	return Result[T]{ok: ok}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) IsOk() bool {
	return !r.IsErr()
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Ok() T {
	return r.ok
}
