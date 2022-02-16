package result

import (
	"fmt"
)

type Result[T any] struct {
	ok  T
	err error
}

func Ok[T any](ok T) Result[T] {
	return Result[T]{ok: ok}
}

func Err[T any](err any) Result[T] {
	return Result[T]{err: newAnyError(err)}
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

func (r Result[T]) ErrVal() any {
	if r.IsErr() {
		return nil
	}
	if ev, _ := r.err.(*errorWithVal); ev != nil {
		return ev.val
	}
	return r.err.Error()
}

func (r Result[T]) Ok() T {
	return r.ok
}

type errorWithVal struct {
	val any
}

func newAnyError(val any) error {
	if err, _ := val.(error); err != nil {
		return err
	}
	return &errorWithVal{val: val}
}

func (a *errorWithVal) Error() string {
	return fmt.Sprintf("%v", a.val)
}
