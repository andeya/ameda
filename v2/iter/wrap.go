package iter

import (
	"github.com/henrylee2cn/ameda/v2/result"
)

func ToMap[T any, B any](iter Iterator[T], f func(T) B) *Map[B] {
	return nil
}

type S struct {
}

type S2 struct {
}

func (s S) Name() {

}
func (s S2) Name() {

}

type X interface {
	S2 | S
	Name()
}

func F[T X](x T) {
	x.Name()
}

func TryFold[T any, B any](iter Iterator[T], init B, f func(B, T) result.Result[B]) result.Result[B] {
	var accum = result.Ok(init)
	for {
		x := iter.Next()
		if x.IsNone() {
			break
		}
		accum = f(accum.Ok(), x.Some())
		if accum.IsErr() {
			return accum
		}
	}
	return accum
}

func TryRFold[T any, B any](iter Iterator[T], init B, f func(B, T) result.Result[B]) result.Result[B] {
	// FIXME
	var accum = result.Ok(init)
	for {
		x := iter.Next()
		if x.IsNone() {
			break
		}
		accum = f(accum.Ok(), x.Some())
		if accum.IsErr() {
			return accum
		}
	}
	return accum
}
