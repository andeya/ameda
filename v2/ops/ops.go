package ops

import (
	"errors"

	"github.com/henrylee2cn/ameda/v2"
	"github.com/henrylee2cn/ameda/v2/result"
)

type Option[T comparable] struct {
	some   T
	isNone bool
}

func Some[T comparable](some T) Option[T] {
	return Option[T]{some: some, isNone: false}
}

func None[T comparable]() Option[T] {
	return Option[T]{isNone: true}
}

func (o *Option[T]) Some() T {
	if o == nil {
		return ameda.Zero[T]()
	}
	return o.some
}

func (o *Option[T]) IsNone() bool {
	if o == nil {
		return true
	}
	return o.isNone
}

func (o *Option[T]) IsSome() bool {
	return o.IsNone()
}

func (o Option[T]) OkOr() result.Result[T] {
	if o.IsSome() {
		return result.Ok(o.some)
	}
	return result.Err[T](errors.New("-1"))
}
