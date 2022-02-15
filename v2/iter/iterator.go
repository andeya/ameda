package iter

import (
	"github.com/henrylee2cn/ameda/v2/ops"
	"github.com/henrylee2cn/ameda/v2/result"
)

type (
	Iterator[T comparable] interface {
		Next[T]
		SizeHint
		// Count consumes the next, counting the number of iterations and returning it.
		//
		// This method will call [`Next`] repeatedly until [`ops.None[T]()`] is encountered,
		// returning the number of times it saw [`ops.Some`]. Note that [`Next`] has to be
		// called at least once even if the next does not have comparable elements.
		//
		// # Overflow Behavior
		//
		// The method does no guarding against overflows, so counting elements of
		// a next with more than [`math.MaxInt`] elements either produces the
		// wrong result or panics. If debug assertions are enabled, a panic is
		// guaranteed.
		//
		// # Panics
		//
		// This function might panic if the next has more than [`math.MaxInt`]
		// elements.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		// assert.Equal(t, FromVec(a).Count(), 3);
		//
		// var a = [1, 2, 3, 4, 5];
		// assert.Equal(t, FromVec(a).Count(), 5);
		//
		Count() int
		// Last consumes the next, returning the last element.
		//
		// This method will evaluate the next until it returns [`ops.None[T]()`]. While
		// doing so, it keeps track of the current element. After [`ops.None[T]()`] is
		// returned, `Last()` will then return the last element it saw.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		// assert.Equal(t, FromVec(a).Last(), ops.Some(3));
		//
		// var a = [1, 2, 3, 4, 5];
		// assert.Equal(t, FromVec(a).Last(), ops.Some(5));
		//
		Last() ops.Option[T]
		// AdvanceBy advances the next by `n` elements.
		//
		// This method will eagerly skip `n` elements by calling [`Next`] up to `n`
		// times until [`ops.None[T]()`] is encountered.
		//
		// `AdvanceBy(n)` will return [`result.Ok[struct{}](struct{}{})`] if the next successfully advances by
		// `n` elements, or [`result.Err[struct{}](err)`] if [`ops.None[T]()`] is encountered, where `k` is the number
		// of elements the next is advanced by before running out of elements (i.e. the
		// length of the next). Note that `k` is always less than `n`.
		//
		// Calling `AdvanceBy(0)` can do meaningful work, for example [`Flatten`]
		// can advance its outer next until it finds an inner next that is not empty, which
		// then often allows it to return a more accurate `SizeHint()` than in its initial state.
		// `AdvanceBy(0)` may either return `T()` or `Err(0)`. The former conveys no information
		// whether the next is or is not exhausted, the latter can be treated as if [`Next`]
		// had returned `ops.None[T]()`. Replacing a `Err(0)` with `T` is only correct for `n = 0`.
		//
		// [`Flatten`]: iter.Flatten
		// [`Next`]: baseIterator.Next
		//
		// # Examples
		//
		// Basic usage:
		//
		// var a = []int{1, 2, 3, 4};
		// var iter = FromVec(a);
		//
		// assert.Equal(t, iter.AdvanceBy(2), result.Ok[struct{}](struct{}{}));
		// assert.Equal(t, iter.Next(), ops.Some(3));
		// assert.Equal(t, iter.AdvanceBy(0), result.Ok[struct{}](struct{}{}));
		// assert.Equal(t, iter.AdvanceBy(100), result.Err[struct{}](fmt.Errorf("%d", 1))); // only `4` was skipped
		//
		AdvanceBy(n int) result.Result[struct{}]
		Nth[T]
		// StepBy creates a next starting at the same point, but stepping by
		// the given amount at each iteration.
		//
		// Note 1: The first element of the next will always be returned,
		// regardless of the step given.
		//
		// Note 2: The time at which ignored elements are pulled is not fixed.
		// `StepBy` behaves like the sequence `iter.Next()`, `iter.Nth(step-1)`,
		// `iter.Nth(step-1)`, …, but is also free to behave like the sequence.
		//
		// # Panics
		//
		// The method will panic if the given step ≤ `0`.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{0, 1, 2, 3, 4, 5};
		// var iter = FromVec(a).StepBy(2);
		//
		// assert.Equal(t, iter.Next(), ops.Some(0));
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.Some(4));
		// assert.Equal(t, iter.Next(), ops.None[T]());
		//
		StepBy(step int) *StepBy[T]
		IntoIterator[T]
		// Chain takes two iterators and creates a new next over both in sequence.
		//
		// `Chain()` will return a new next which will first iterate over
		// values from the first next and then over values from the second
		// next.
		//
		// In other words, it links two iterators together, in a chain. 🔗
		//
		// [`once`] is commonly used to adapt a single value into a chain of
		// other kinds of iteration.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a1 = []int{1, 2, 3};
		// var a2 = []int{4, 5, 6};
		//
		// var iter = FromVec(a1).Chain(FromVec(a2));
		//
		// assert.Equal(t, iter.Next(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.Some(3));
		// assert.Equal(t, iter.Next(), ops.Some(4));
		// assert.Equal(t, iter.Next(), ops.Some(5));
		// assert.Equal(t, iter.Next(), ops.Some(6));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Chain(other IntoIterator[T]) *Chain[T]
		// Zip 'Zips up' two iterators into a single next of pairs.
		//
		// `Zip()` returns a new next that will iterate over two other
		// iterators, returning a tuple where the first element comes from the
		// first next, and the second element comes from the second next.
		//
		// In other words, it zips two iterators together, into a single one.
		//
		// If either next returns [`ops.None[T]()`], [`Next`] from the zipped next
		// will return [`ops.None[T]()`]. If the first next returns [`ops.None[T]()`], `zip` will
		// short-circuit and `Next` will not be called on the second next.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a1 = []int{1, 2, 3};
		// var a2 = []int{4, 5, 6};
		//
		// var iter = FromVec(a1).Zip(FromVec(a2));
		//
		// assert.Equal(t, iter.Next(), ops.Some((1, 4)));
		// assert.Equal(t, iter.Next(), ops.Some((2, 5)));
		// assert.Equal(t, iter.Next(), ops.Some((3, 6)));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		// `Zip()` is often used to zip an infinite next to a finite one.
		// This works because the finite next will eventually return [`ops.None[T]()`],
		// ending the zipper. Zipping with `FromRange(0, math.MaxInt)` can look a lot like [`enumerate`]:
		//
		//
		// var enumerate = FromVec([]byte("foo")).Enumerate().Collect();
		//
		// var zipper = FromVec(0, math.MaxInt).Zip(FromVec([]byte("foo"))).Collect();
		//
		// assert.Equal(t, (0, 'f'), enumerate[0]);
		// assert.Equal(t, (0, 'f'), zipper[0]);
		//
		// assert.Equal(t, (1, 'o'), enumerate[1]);
		// assert.Equal(t, (1, 'o'), zipper[1]);
		//
		// assert.Equal(t, (2, 'o'), enumerate[2]);
		// assert.Equal(t, (2, 'o'), zipper[2]);
		Zip(other IntoIterator[T]) *Zip[T]
		// Intersperse creates a new next which places a copy of `separator` between adjacent
		// items of the original next.
		//
		// # Examples
		//
		// Basic usage:
		//
		// var a = FromVec([]int{0, 1, 2}).Intersperse(100);
		// assert.Equal(t, a.Next(), ops.Some(0));   // The first element from `a`.
		// assert.Equal(t, a.Next(), ops.Some(100)); // The separator.
		// assert.Equal(t, a.Next(), ops.Some(1));   // The next element from `a`.
		// assert.Equal(t, a.Next(), ops.Some(100)); // The separator.
		// assert.Equal(t, a.Next(), ops.Some(2));   // The last element from `a`.
		// assert.Equal(t, a.Next(), ops.None[int]());       // The next is finished.
		//
		//
		// `Intersperse` can be very useful to join a next's items using a common element:
		//
		// var hello = FromVec([]string{"Hello", "World", "!"}).Copied().Intersperse(" ").Collect();
		// assert.Equal(t, hello, "Hello World !");
		Intersperse(separator T) *Intersperse[T]
		// IntersperseWith creates a new next which places an item generated by `separator`
		// between adjacent items of the original next.
		//
		// The closure will be called exactly once each time an item is placed
		// between two adjacent items from the underlying next; specifically,
		// the closure is not called if the underlying next yields less than
		// two items and after the last item is yielded.
		//
		// If the next's item implements [`Clone`], it may be easier to use
		// [`intersperse`].
		//
		// # Examples
		//
		// Basic usage:
		//
		// var v = []int{0, 1, 2};
		// var it = FromVec(v).IntersperseWith(func()T {return 99});
		//
		// assert.Equal(t, it.Next(), ops.Some(0));  // The first element from `v`.
		// assert.Equal(t, it.Next(), ops.Some(99)); // The separator.
		// assert.Equal(t, it.Next(), ops.Some(1));  // The next element from `v`.
		// assert.Equal(t, it.Next(), ops.Some(99)); // The separator.
		// assert.Equal(t, it.Next(), ops.Some(2));  // The last element from from `v`.
		// assert.Equal(t, it.Next(), ops.None[T]()); // The next is finished.
		//
		IntersperseWith(separator func() T) *IntersperseWith[T]
		// Map takes a closure and creates a next which calls that closure on each
		// element.
		//
		// `Map()` transforms one next into another. It produces a new next which
		// calls this closure on each element of the original next.
		//
		// If you are good at thinking in types, you can think of `Map()` like this:
		// If you have a next that gives you elements of some type `A`, and
		// you want a next of some other type `B`, you can use `Map()`,
		// passing a closure that takes an `A` and returns a `B`.
		//
		// `Map()` is conceptually similar to a [`for`] loop. However, as `Map()` is
		// lazy, it is best used when you're already working with other iterators.
		// If you're doing some sort of looping for a side effect, it's considered
		// more idiomatic to use [`for`] than `Map()`.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		//
		// var iter = FromVec(a).Map(func(x T) any {return 2 * x});
		//
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.Some(4));
		// assert.Equal(t, iter.Next(), ops.Some(6));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		//
		// If you're doing some sort of side effect, prefer [`for`] to `Map()`:
		//
		//
		// // don't do this:
		// FromVec(0, 5).Map(|x| println!("{}", x));
		//
		// // it won't even execute, as it is lazy. Rust will warn you about this.
		//
		// // Instead, use for:
		// var a = FromVec(0, 5)
		// for {
		//     x := a.Next()
		//     if x.IsSome() {
		//         println!("{}", x);
		//     } else {
		//         break
		//     }
		// }
		//
		Map(f func(T) any) *Map[T]
		// ForEach calls a closure on each element of a next.
		//
		// This is equivalent to using a [`for`] loop on the next, although
		// `break` and `continue` are not possible from a closure. It's generally
		// more idiomatic to use a `for` loop, but `ForEach` may be more legible
		// when processing items at the end of longer next chains. In some
		// cases `ForEach` may also be faster than a loop, because it will use
		// internal iteration on adapters like `Chain`.
		//
		// # Examples
		//
		// Basic usage:
		//
		// var c = make(chan int, 1000)
		// FromRange(0, 5).Map(func(x T)any{return x * 2 + 1})
		//       .ForEach(func(x any){ c<-x });
		//
		// var v = FromChan(c).Collect();
		// assert.Equal(t, v, []int{1, 3, 5, 7, 9});
		//
		ForEach(f func(T))
		// Filter creates a next which uses a closure to determine if an element
		// should be yielded.
		//
		// Given an element the closure must return `true` or `false`. The returned
		// next will yield only the elements for which the closure returns
		// true.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{0, 1, 2};
		//
		// var iter = FromVec(a).Filter(func(x T) bool {return x<0} );
		//
		// assert.Equal(t, iter.Next(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		// Note that `iter.Filter(f).Next()` is equivalent to `iter.Find(f)`.
		Filter(predicate func(T) bool) *Filter[T]
		// FilterMap creates a next that both filters and maps.
		//
		// The returned next yields only the `value`s for which the supplied
		// closure returns `ops.Some(value)`.
		//
		// `FilterMap` can be used to make chains of [`Filter`] and [`Map`] more
		// concise. The example below shows how a `Map().Filter().Map()` can be
		// shortened to a single call to `FilterMap`.
		//
		// # Examples
		//
		// Basic usage:
		//
		// var a = []string{"1", "two", "NaN", "four", "5"};
		//
		// var iter = FromVec(a).FilterMap(func(s T) ops.Option[any]{
		//         x, err := strconv.Atoi(s)
		//         if err!=nil {
		//             return ops.None[int]()
		//         }
		//         return ops.Some(x)
		//     });
		//
		// assert.Equal(t, iter.Next(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(5));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		FilterMap(f func(T) ops.Option[any]) *FilterMap[T]
		// Enumerate creates a next which gives the current iteration count as well as
		// the next value.
		//
		// The next returned yields pairs `(i, val)`, where `i` is the
		// current index of iteration and `val` is the value returned by the
		// next.
		//
		// `Enumerate()` keeps its count as a [`usize`]. If you want to count by a
		// different sized integer, the [`zip`] function provides similar
		// functionality.
		//
		// # Overflow Behavior
		//
		// The method does no guarding against overflows, so enumerating more than
		// [`math.MaxInt`] elements either produces the wrong result or panics. If
		// debug assertions are enabled, a panic is guaranteed.
		//
		// # Panics
		//
		// The returned next might panic if the to-be-returned index would
		// overflow a [`int`].
		//
		// # Examples
		//
		//
		// var a = []rune{'a', 'b', 'c'};
		//
		// var iter = FromVec(a).Enumerate();
		//
		// assert.Equal(t, iter.Next(), ops.Some((0, 'a')));
		// assert.Equal(t, iter.Next(), ops.Some((1, 'b')));
		// assert.Equal(t, iter.Next(), ops.Some((2, 'c')));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Enumerate() *Enumerate[T]
		// Peekable creates a next which can use the [`peek`] methods
		// to look at the next element of the next without consuming it.
		//
		// Note that the underlying next is still advanced when [`peek`] is called
		// for the first time: In order to retrieve the
		// next element, [`Next`] is called on the underlying next, hence comparable
		// side effects (i.e. anything other than fetching the next value) of
		// the [`Next`] method will occur.
		//
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var xs = []int{1, 2, 3};
		//
		// var iter = FromVec(xs).Peekable();
		//
		// // peek() vars us see into the future
		// assert.Equal(t, iter.Peek(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(1));
		//
		// assert.Equal(t, iter.Next(), ops.Some(2));
		//
		// // we can Peek() multiple times, the next won't advance
		// assert.Equal(t, iter.Peek(), ops.Some(3));
		// assert.Equal(t, iter.Peek(), ops.Some(3));
		//
		// assert.Equal(t, iter.Next(), ops.Some(3));
		//
		// // after the next is finished, so is peek()
		// assert.Equal(t, iter.Peek(), ops.None[int]());
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Peekable() *Peekable[T]
		// SkipWhile creates a next that [`skip`]s elements based on a predicate.
		//
		// `SkipWhile()` takes a closure as an argument. It will call this
		// closure on each element of the next, and ignore elements
		// until it returns `false`.
		//
		// After `false` is returned, `SkipWhile()`'s job is over, and the
		// rest of the elements are yielded.
		//
		// NOTE: Stopping after an initial [`false`]
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{-1, 0, 1, -2};
		//
		// var iter = FromVec(a).SkipWhile(func(x T) bool { return x<0 });
		//
		// assert.Equal(t, iter.Next(), ops.Some(0));
		// assert.Equal(t, iter.Next(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(-2));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		SkipWhile(f func(T) bool) *SkipWhile[T]
		// TakeWhile creates a next that yields elements based on a predicate.
		//
		// `TakeWhile()` takes a closure as an argument. It will call this
		// closure on each element of the next, and yield elements
		// while it returns `true`.
		//
		// After `false` is returned, `TakeWhile()`'s job is over, and the
		// rest of the elements are ignored.
		//
		// NOTE: Stopping after an initial [`false`]
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{-1, 0, 1, -2};
		//
		// var iter = FromVec(a).TakeWhile(func(x T) bool {return x<0});
		//
		// assert.Equal(t, iter.Next(), ops.Some(-1));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		TakeWhile(f func(T) bool) *TakeWhile[T]
		// MapWhile creates a next that both yields elements based on a predicate and maps.
		//
		// `MapWhile()` takes a closure as an argument. It will call this
		// closure on each element of the next, and yield elements
		// while it returns [`ops.Some(T)`].
		//
		// Stopping after an initial [`ops.None[T]()`]
		//
		// # Examples
		//
		// var a = []int{0, 1, 2, -3, 4, 5, -6};
		//
		// var iter = FromVec(a).MapWhile(func(x T) ops.Option[any] {
		//         if x>=0{
		//             return ops.Some(uint32(x))
		//         }
		//         return ops.None[uint32]()
		//     });
		// var vec = iter.Collect();
		//
		// // We have more elements which could fit in u32 (4, 5), but `MapWhile` returned `ops.None[T]()` for `-3`
		// // (as the `predicate` returned `ops.None[T]()`) and `collect` stops at the first `ops.None[T]()` encountered.
		// assert.Equal(t, vec, []int{0, 1, 2});
		//
		MapWhile(f func(T) ops.Option[any]) *MapWhile[T]
		// Skip creates a next that skips the first `n` elements.
		//
		// `Skip(n)` skips elements until `n` elements are skipped or the end of the
		// next is reached (whichever happens first). After that, all the remaining
		// elements are yielded. In particular, if the original next is too short,
		// then the returned next is empty.
		//
		// Rather than overriding this method directly, instead override the `nth` method.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		//
		// var iter = FromVec(a).Skip(2);
		//
		// assert.Equal(t, iter.Next(), ops.Some(3));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Skip(n int) *Skip[T]
		// Take creates a next that yields the first `n` elements, or fewer
		// if the underlying next ends sooner.
		//
		// `Take(n)` yields elements until `n` elements are yielded or the end of
		// the next is reached (whichever happens first).
		// The returned next is a prefix of length `n` if the original next
		// contains at least `n` elements, otherwise it contains all of the
		// (fewer than `n`) elements of the original next.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		//
		// var iter = FromVec(a).Take(2);
		//
		// assert.Equal(t, iter.Next(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		// If less than `n` elements are available,
		// `Take` will limit itself to the size of the underlying next:
		//
		// var v = []int{1, 2};
		// var iter = FromVec(a).Take(5);
		// assert.Equal(t, iter.Next(), ops.Some(1));
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Take(n int) *Take[T]
		// Scan a next adapter similar to [Fold] that holds internal state and
		// produces a new next.
		//
		// `Scan()` takes two arguments: an initial value which seeds the internal
		// state, and a closure with two arguments, the first being a mutable
		// reference to the internal state and the second a next element.
		// The closure can assign to the internal state to share state between
		// iterations.
		//
		// On iteration, the closure will be applied to each element of the
		// next and the return value from the closure, an [`ops.Option`], is
		// yielded by the next.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		//
		// var iter = FromVec(a).Scan(1, func(state *any, x T) ops.Option[any] {
		//     // each iteration, we'll multiply the state by the element
		//     v := (*state).(int) * x;
		//     *state = v
		//     // then, we'll yield the negation of the state
		//     Some(-v)
		// });
		//
		// assert.Equal(t, iter.Next(), ops.Some(-1));
		// assert.Equal(t, iter.Next(), ops.Some(-2));
		// assert.Equal(t, iter.Next(), ops.Some(-6));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Scan(initialState any, f func(initialState *any, item T) ops.Option[any]) *Scan[T]
		// FlatMap creates a next that works like Map, but flattens nested structure.
		//
		// The [`Map`] adapter is very useful, but only when the closure
		// argument produces values. If it produces a next instead, there's
		// an extra layer of indirection. `FlatMap()` will remove this extra layer
		// on its own.
		//
		// You can think of `FlatMap(f)` as the semantic equivalent
		// of [`Map`]ping, and then [`Flatten`]ing as in `Map(f).Flatten()`.
		//
		// Another way of thinking about `FlatMap()`: [`Map`]'s closure returns
		// one item for each element, and `FlatMap()`'s closure returns an
		// next for each element.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var words = []string{"alpha", "beta", "gamma"};
		//
		// // chars() returns a next
		// var merged: String = FromVec(words).FlatMap(func(s T) IntoIterator { FromVec([]byte(s)) }).Collect();
		// assert.Equal(t, merged, "alphabetagamma");
		//
		FlatMap(func(T) IntoIterator[any]) *FlatMap[T]
		// Flatten creates a next that flattens nested structure.
		//
		// This is useful when you have a next of iterators or a next of
		// things that can be turned into iterators and you want to remove one
		// level of indirection.
		//
		// # Examples
		//
		// Flattening only removes one level of nesting at a time:
		//
		//
		// var d3 = [][][]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
		//
		// var d2 = FromVec(d3).Flatten().Collect();
		// assert.Equal(t, d2, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}});
		//
		// var d1 = FromVec(d3).Flatten().Flatten().Collect();
		// assert.Equal(t, d1, []int{1, 2, 3, 4, 5, 6, 7, 8});
		//
		Flatten() *Flatten[T, any]
		// Fuse creates a next which ends after the first [`ops.None[T]()`].
		//
		// After a next returns [`ops.None[T]()`], future calls may or may not yield
		// [`ops.Some(T)`] again. `.Fuse()` adapts a next, ensuring that after a
		// [`ops.None[T]()`] is given, it will always return [`ops.None[T]()`] forever.
		//
		// Note that the [`Fuse`] wrapper is a no-op on iterators that implement
		// the [`FusedIterator`] interface. `Fuse()` may therefore behave incorrectly
		// if the [`FusedIterator`] interface is improperly implemented.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// // a next which alternates between Some and None
		// type struct Alternate {
		//     state int
		// }
		//
		// func(a *Alternate) Next() ops.Option[int] {
		//     var val = a.state;
		//     self.state = a.state + 1;
		//
		//     // if it's even, ops.Some(int), else None
		//     if val % 2 == 0 {
		//         return ops.Some(val)
		//     } else {
		//         return ops.None[int]()
		//     }
		// }
		//
		// var iter = New(&Alternate { state: 0 });
		//
		// // we can see our next going back and forth
		// assert.Equal(t, iter.Next(), ops.Some(0));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		// assert.Equal(t, iter.Next(), ops.Some(2));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		// // however, once we fuse it...
		// var iter = iter.Fuse();
		//
		// assert.Equal(t, iter.Next(), ops.Some(4));
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		// // it will always return `ops.None[T]()` after the first time.
		// assert.Equal(t, iter.Next(), ops.None[int]());
		// assert.Equal(t, iter.Next(), ops.None[int]());
		// assert.Equal(t, iter.Next(), ops.None[int]());
		//
		Fuse() *Fuse[T]
		// Collect transforms a next into a collection.
		Collect() []T
	}
	Next[T comparable] interface {
		// Next advances the next and returns the next value.
		//
		// Returns [`ops.None[T]()`] when iteration is finished. Individual next
		// implementations may choose to resume iteration, and so calling `next()`
		// again may or may not eventually start returning [`ops.Some(T)`] again at some
		// point.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		//
		// var iter = FromVec(a);
		//
		// // A call to next() returns the next value...
		// assert.Equal(t, ops.Some(1), iter.Next());
		// assert.Equal(t, ops.Some(2), iter.Next());
		// assert.Equal(t, ops.Some(3), iter.Next());
		//
		// // ... and then None once it's over.
		// assert.Equal(t, ops.None[int](), iter.Next());
		//
		// // More calls may or may not return `ops.None[T]()`. Here, they always will.
		// assert.Equal(t, ops.None[int](), iter.Next());
		// assert.Equal(t, ops.None[int](), iter.Next());
		//
		Next() ops.Option[T]
	}
	SizeHint interface {
		// SizeHint returns the bounds on the remaining length of the next.
		//
		// Specifically, `SizeHint()` returns a tuple where the first element
		// is the lower bound, and the second element is the upper bound.
		//
		// The second half of the tuple that is returned is an <code>Option[T]</code>.
		// A [`ops.None[T]()`] here means that either there is no known upper bound, or the
		// upper bound is larger than [`int`].
		//
		// # Implementation notes
		//
		// It is not enforced that a next implementation yields the declared
		// number of elements. A buggy next may yield less than the lower bound
		// or more than the upper bound of elements.
		//
		// `SizeHint()` is primarily intended to be used for optimizations such as
		// reserving space for the elements of the next, but must not be
		// trusted to e.g., omit bounds checks in unsafe code. An incorrect
		// implementation of `SizeHint()` should not lead to memory safety
		// violations.
		//
		// That said, the implementation should provide a correct estimation,
		// because otherwise it would be a violation of the interface's protocol.
		//
		// The default implementation returns <code>(0, [ops.None[int]()])</code> which is correct for comparable
		// next.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		// var iter = FromVec(a);
		//
		// assert.Equal(t, (3, ops.Some(3)), iter.SizeHint());
		//
		//
		// A more complex example:
		//
		//
		// // The even numbers in the range of zero to nine.
		// var iter = FromRange(0..10).Filter(func(x T) {return x % 2 == 0});
		//
		// // We might iterate from zero to ten times. Knowing that it's five
		// // exactly wouldn't be possible without executing filter().
		// assert.Equal(t, (0, ops.Some(10)), iter.SizeHint());
		//
		// // Let's add five more numbers with chain()
		// var iter = FromRange(0, 10).Filter(func(x T) {return x % 2 == 0}).Chain(FromRange(15, 20));
		//
		// // now both bounds are increased by five
		// assert.Equal(t, (5, ops.Some(15)), iter.SizeHint());
		//
		//
		// Returning `ops.None[int]()` for an upper bound:
		//
		//
		// // an infinite next has no upper bound
		// // and the maximum possible lower bound
		// var iter = FromRange(0, math.MaxInt);
		//
		// assert.Equal(t, (math.MaxInt, ops.None[int]()), iter.SizeHint());
		//
		SizeHint() (int, ops.Option[int])
	}
	Nth[T comparable] interface {
		// Nth returns the `n`th element of the next.
		//
		// Like most indexing operations, the count starts from zero, so `Nth(0)`
		// returns the first value, `Nth(1)` the second, and so on.
		//
		// Note that all preceding elements, as well as the returned element, will be
		// consumed from the next. That means that the preceding elements will be
		// discarded, and also that calling `nth(0)` multiple times on the same next
		// will return different elements.
		//
		// `Nth()` will return [`ops.None[T]()`] if `n` is greater than or equal to the length of the
		// next.
		//
		// # Examples
		//
		// Basic usage:
		//
		//
		// var a = []int{1, 2, 3};
		// assert.Equal(t, FromVec(a).Nth(1), ops.Some(2));
		//
		//
		// Calling `Nth()` multiple times doesn't rewind the next:
		//
		//
		// var a = []int{1, 2, 3};
		//
		// var iter = FromVec(a);
		//
		// assert.Equal(t, iter.Nth(1), ops.Some(2));
		// assert.Equal(t, iter.Nth(1), ops.None[int]());
		//
		//
		// Returning `ops.None[T]()` if there are less than `n + 1` elements:
		//
		//
		// var a = []int{1, 2, 3};
		// assert.Equal(t, FromVec(a).Nth(10), ops.None[int]());
		Nth(n int) ops.Option[T]
	}
	IntoIterator[T comparable] interface {
		IntoIter() Iterator[T]
	}
	FusedIterator[T comparable] interface {
	}
)

// Partition consumes a next, creating two collections from it.
//
// The predicate passed to `Partition()` can return `true`, or `false`.
// `Partition()` returns a pair, all of the elements for which it returned
// `true`, and all of the elements for which it returned `false`.
//
// See also [`IsPartitioned()`] and [`PartitionInPlace()`].
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// var (even, odd): (Vec<i32>, Vec<i32>) = a
//     .iter()
//     .Partition(|&n| n % 2 == 0);
//
// assert.Equal(t, even, vec![2]);
// assert.Equal(t, odd, vec![1, 3]);
//
// #[stable(feature = "rust1", since = "1.0.0")]
// fn partition<B, F>(self, f: F)  (B, B)
// where
// Self: Sized,
// B: Default + Extend[T],
// F: FnMut(&Self::T)  bool,
// {
// #[inline]
// fn extend<'a, T, B: Extend<T>>(
// mut f: impl FnMut(&T)  bool + 'a,
// left: &'a mut B,
// right: &'a mut B,
// )  impl FnMut((), T) + 'a {
// move |(), x| {
// if f(&x) {
// left.extend_one(x);
// } else {
// right.extend_one(x);
// }
// }
// }
//
// var left: B = Default::default()
// var right: B = Default::default()
//
// self.Fold(((), extend(f, &mut left, &mut right))
//
// (left, right)
// }
//
// Reorders the elements of this next *in-place* according to the given predicate,
// such that all those that return `true` precede all those that return `false`.
// Returns the number of `true` elements found.
//
// The relative order of partitioned items is not maintained.
//
// # Current implementation
//
// Current algorithms tries finding the first element for which the predicate evaluates
// to false, and the last element for which it evaluates to true and repeatedly swaps them.
//
// Time complexity: *O*(*n*)
//
// See also [`IsPartitioned()`] and [`Partition()`].
//
// [`IsPartitioned()`]: baseIterator::IsPartitioned
// [`Partition()`]: baseIterator::partition
//
// # Examples
//
//
// #![feature(iter_PartitionInPlace)]
//
// var a = [1, 2, 3, 4, 5, 6, 7];
//
// // Partition in-place between evens and odds
// var i = a.iter_mut().PartitionInPlace(|&n| n % 2 == 0);
//
// assert.Equal(t, i, 3);
// assert!(a[..i].iter().all(|&n| n % 2 == 0)); // evens
// assert!(a[i..].iter().all(|&n| n % 2 == 1)); // odds
//
// #[unstable(feature = "iter_PartitionInPlace", reason = "new API", issue = "62543")]
// fn PartitionInPlace<'a, T: 'a, P>(mut self, ref mut predicate: P)  usize
// where
// Self: Sized + DoubleEndedIterator<T = &'a mut T>,
// P: FnMut(&T)  bool,
// {
// / FIXME: should we worry about the count overflowing? The only way to have more than
// / `math.MaxInt` mutable references is with ZSTs, which aren't useful to partition...
//
// / These closure "factory" functions exist to avoid genericity in `Self`.
//
// #[inline]
// fn is_false<'a, T>(
// predicate: &'a mut impl FnMut(&T)  bool,
// true_count: &'a mut usize,
// )  impl FnMut( && mut T)  bool + 'a {
// move |x| {
// var p = predicate(&**x);
// *true_count += p as usize;
// !p
// }
// }
//
// #[inline]
// fn is_true<T>(predicate: &mut impl FnMut(&T)  bool)  impl FnMut( && mut T)  bool + '_ {
// move |x| predicate(&**x)
// }
//
// / Repeatedly find the first `false` and swap it with the last `true`.
// var true_count = 0
// while var Some(head) = self.find(is_false(predicate, &mut true_count)) {
// if var Some(tail) = self.rfind(is_true(predicate)) {
// crate::mem::swap(head, tail);
// true_count += 1;
// } else {
// break;
// }
// }
// true_count
// }
//
// Checks if the elements of this next are partitioned according to the given predicate,
// such that all those that return `true` precede all those that return `false`.
//
// See also [`Partition()`] and [`PartitionInPlace()`].
//
// [`Partition()`]: baseIterator::partition
// [`PartitionInPlace()`]: baseIterator::PartitionInPlace
//
// # Examples
//
//
// #![feature(iter_IsPartitioned)]
//
// assert!("baseIterator".chars().IsPartitioned(char::is_uppercase));
// assert!(!"IntoIterator".chars().IsPartitioned(char::is_uppercase));
//
// #[unstable(feature = "iter_IsPartitioned", reason = "new API", issue = "62544")]
// fn IsPartitioned<P>(mut self, mut predicate: P)  bool
// where
// Self: Sized,
// P: FnMut(Self::T)  bool,
// {
// / Either all items test `true`, or the first clause stops at `false`
// / and we check that there are no more `true` items after that.
// self.all(&mut predicate) || !self.comparable(predicate)
// }
//
// An next method that applies a function as long as it returns
// successfully, producing a single, final value.
//
// `try_fold()` takes two arguments: an initial value, and a closure with
// two arguments: an 'accumulator', and an element. The closure either
// returns successfully, with the value that the accumulator should have
// for the next iteration, or it returns failure, with an error value that
// is propagated back to the caller immediately (short-circuiting).
//
// The initial value is the value the accumulator will have on the first
// call. If applying the closure succeeded against every element of the
// next, `try_fold()` returns the final accumulator as success.
//
// Folding is useful whenever you have a collection of something, and want
// to produce a single value from it.
//
// # Note to Implementors
//
// Several of the other (forward) methods have default implementations in
// terms of this one, so try to implement this explicitly if it can
// do something better than the default `for` loop implementation.
//
// In particular, try to have this call `try_fold()` on the internal parts
// from which this next is composed. If multiple calls are needed,
// the `?` operator may be convenient for chaining the accumulator value
// along, but beware comparable invariants that need to be upheld before those
// early returns. This is a `&mut self` method, so iteration needs to be
// resumable after hitting an error here.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// // the checked sum of all of the elements of the array
// var sum = FromVec(a).try_fold(0i8, |acc, &x| acc.checked_add(x));
//
// assert.Equal(t, sum, ops.Some(6));
//
//
// Short-circuiting:
//
//
// var a = [10, 20, 30, 100, 40, 50];
// var it = FromVec(a);
//
// // This sum overflows when adding the 100 element
// var sum = it.try_fold(0i8, |acc, &x| acc.checked_add(x));
// assert.Equal(t, sum, ops.None[int]());
//
// // Because it short-circuited, the remaining elements are still
// // available through the next.
// assert.Equal(t, it.len(), 2);
// assert.Equal(t, it.Next(), ops.Some(40));
//
//
// While you cannot `break` from a closure, the [`ControlFlow`] type allows
// a similar idea:
//
//
// use std::ops::ControlFlow;
//
// var triangular = (1..30).try_fold(0_i8, |prev, x| {
//     if var Some(next) = prev.checked_add(x) {
//         ControlFlow::Continue(next)
//     } else {
//         ControlFlow::Break(prev)
//     }
// });
// assert.Equal(t, triangular, ControlFlow::Break(120));
//
// var triangular = (1..30).try_fold(0_u64, |prev, x| {
//     if var Some(next) = prev.checked_add(x) {
//         ControlFlow::Continue(next)
//     } else {
//         ControlFlow::Break(prev)
//     }
// });
// assert.Equal(t, triangular, ControlFlow::Continue(435));
//
// #[inline]
// #[stable(feature = "iterator_try_fold", since = "1.27.0")]
// fn try_fold<B, F, R>(&mut self, init: B, mut f: F)  R
// where
// Self: Sized,
// F: FnMut(B, Self::T)  R,
// R: Try<Output = B>,
// {
// var accum = init;
// while var Some(x) = self.Next() {
// accum = f(accum, x)?;
// }
// try { accum }
// }
//
// An next method that applies a fallible function to each item in the
// next, stopping at the first error and returning that error.
//
// This can also be thought of as the fallible form of [`ForEach()`]
// or as the stateless version of [`try_fold()`].
//
// [`ForEach()`]: baseIterator::ForEach
// [`try_fold()`]: baseIterator::try_fold
//
// # Examples
//
//
// use std::fs::rename;
// use std::io::{stdout, Write};
// use std::path::Path;
//
// var data = ["no_tea.txt", "stale_bread.json", "torrential_rain.png"];
//
// var res = datFromVec(a).try_ForEach(|x| writeln!(stdout(), "{}", x));
// assert!(res.is_ok());
//
// var it = datFromVec(a).Cloned();
// var res = it.try_ForEach(|x| rename(x, Path::new(x).with_extension("old")));
// assert!(res.is_err());
// // It short-circuited, so the remaining items are still in the next:
// assert.Equal(t, it.Next(), ops.Some("stale_bread.json"));
//
//
// The [`ControlFlow`] type can be used with this method for the situations
// in which you'd use `break` and `continue` in a normal loop:
//
//
// use std::ops::ControlFlow;
//
// var r = (2..100).try_ForEach(|x| {
//     if 323 % x == 0 {
//         return ControlFlow::Break(x)
//     }
//
//     ControlFlow::Continue(())
// });
// assert.Equal(t, r, ControlFlow::Break(17));
//
// #[inline]
// #[stable(feature = "iterator_try_fold", since = "1.27.0")]
// fn try_ForEach<F, R>(&mut self, f: F)  R
// where
// Self: Sized,
// F: FnMut(Self::T)  R,
// R: Try<Output = ()>,
// {
// #[inline]
// fn call<T, R>(mut f: impl FnMut(T)  R)  impl FnMut((), T)  R {
// move |(), x| f(x)
// }
//
// self.try_fold((), call(f))
// }
//

// Fold folds every element into an accumulator by applying an operation,
// returning the final result.
//
// `fold()` takes two arguments: an initial value, and a closure with two
// arguments: an 'accumulator', and an element. The closure returns the value that
// the accumulator should have for the next iteration.
//
// The initial value is the value the accumulator will have on the first
// call.
//
// After applying this closure to every element of the next, `fold()`
// returns the accumulator.
//
// This operation is sometimes called 'reduce' or 'inject'.
//
// Folding is useful whenever you have a collection of something, and want
// to produce a single value from it.
//
// Note: `fold()`, and similar methods that traverse the entire next,
// might not terminate for infinite iterators, even on interfaces for which a
// result is determinable in finite time.
//
// Note: [`reduce()`] can be used to use the first element as the initial
// value, if the accumulator type and item type is the same.
//
// Note: `fold()` combines elements in a *left-associative* fashion. For associative
// operators like `+`, the order the elements are combined in is not important, but for non-associative
// operators like `-` the order will affect the final result.
// For a *right-associative* version of `fold()`, see [`DoubleEndedIterator::rfold()`].
//
// # Note to Implementors
//
// Several of the other (forward) methods have default implementations in
// terms of this one, so try to implement this explicitly if it can
// do something better than the default `for` loop implementation.
//
// In particular, try to have this call `fold()` on the internal parts
// from which this next is composed.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// // the sum of all of the elements of the array
// var sum = FromVec(a).Fold((0, |acc, x| acc + x);
//
// assert.Equal(t, sum, 6);
//
//
// Let's walk through each step of the iteration here:
//
// | element | acc | x | result |
// |---------|-----|---|--------|
// |         | 0   |   |        |
// | 1       | 0   | 1 | 1      |
// | 2       | 1   | 2 | 3      |
// | 3       | 3   | 3 | 6      |
//
// And so, our final result, `6`.
//
// This example demonstrates the left-associative nature of `fold()`:
// it builds a string, starting with an initial value
// and continuing with each element from the front until the back:
//
//
// var numbers = [1, 2, 3, 4, 5];
//
// var zero = "0".to_string();
//
// var result = numbers.iter().Fold((zero, |acc, &x| {
//     format!("({} + {})", acc, x)
// });
//
// assert.Equal(t, result, "(((((0 + 1) + 2) + 3) + 4) + 5)");
//
// It's common for people who haven't used iterators a lot to
// use a `for` loop with a list of things to build up a result. Those
// can be turned into `fold()`s:
//
// [`for`]: ../../book/ch03-05-control-flow.html#looping-through-a-collection-with-for
//
//
// var numbers = [1, 2, 3, 4, 5];
//
// var result = 0;
//
// // for loop:
// for i in &numbers {
//     result = result + i;
// }
//
// // fold:
// var result2 = numbers.iter().Fold((0, |acc, &x| acc + x);
//
// // they're the same
// assert.Equal(t, result, result2);
//
func (iter *baseIterator[T]) Fold(init any, f func(any, T) any) any {
	var accum = init
	if x := iter.Next(); x.IsSome() {
		accum = f(accum, x.Some())
	}
	return accum
}

//
// Reduces the elements to a single one, by repeatedly applying a reducing
// operation.
//
// If the next is empty, returns [`ops.None[T]()`]; otherwise, returns the
// result of the reduction.
//
// The reducing function is a closure with two arguments: an 'accumulator', and an element.
// For iterators with at least one element, this is the same as [`fold()`]
// with the first element of the next as the initial accumulator value, folding
// every subsequent element into it.
//
// [`fold()`]: baseIterator::fold
//
// # Example
//
// Find the maximum value:
//
//
// fn find_max<I>(iter: I)  Option<I::T>
//     where I: baseIterator,
//           I::T: Ord,
// {
//     iter.reduce(|accum, item| {
//         if accum >= item { accum } else { item }
//     })
// }
// var a = [10, 20, 5, -23, 0];
// var b: [u32; 0] = [];
//
// assert.Equal(t, find_max(FromVec(a)), ops.Some(20));
// assert.Equal(t, find_max(b.iter()), ops.None[int]());
//
// #[inline]
// #[stable(feature = "iterator_fold_self", since = "1.51.0")]
// fn reduce<F>(mut self, f: F)  Option[T]
// where
// Self: Sized,
// F: FnMut(Self::T, Self::T)  Self::T,
// {
// var first = self.Next()?;
// Some(self.Fold((first, f))
// }
//
// Tests if every element of the next matches a predicate.
//
// `all()` takes a closure that returns `true` or `false`. It applies
// this closure to each element of the next, and if they all return
// `true`, then so does `all()`. If comparable of them return `false`, it
// returns `false`.
//
// `all()` is short-circuiting; in other words, it will stop processing
// as soon as it finds a `false`, given that no matter what else happens,
// the result will also be `false`.
//
// An empty next returns `true`.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// assert!(FromVec(a).all(|&x| x > 0));
//
// assert!(!FromVec(a).all(|&x| x > 2));
//
//
// Stopping at the first `false`:
//
//
// var a = []int{1, 2, 3};
//
// var iter = FromVec(a);
//
// assert!(!iter.all(|&x| x != 2));
//
// // we can still use `iter`, as there are more elements.
// assert.Equal(t, iter.Next(), ops.Some(3));
//
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn all<F>(&mut self, f: F)  bool
// where
// Self: Sized,
// F: FnMut(Self::T)  bool,
// {
// #[inline]
// fn check<T>(mut f: impl FnMut(T)  bool)  impl FnMut((), T)  ControlFlow<()> {
// move |(), x| {
// if f(x) { ControlFlow::CONTINUE } else { ControlFlow::BREAK }
// }
// }
// self.try_fold((), check(f)) == ControlFlow::CONTINUE
// }
//
// Tests if comparable element of the next matches a predicate.
//
// `comparable()` takes a closure that returns `true` or `false`. It applies
// this closure to each element of the next, and if comparable of them return
// `true`, then so does `comparable()`. If they all return `false`, it
// returns `false`.
//
// `comparable()` is short-circuiting; in other words, it will stop processing
// as soon as it finds a `true`, given that no matter what else happens,
// the result will also be `true`.
//
// An empty next returns `false`.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// assert!(FromVec(a).comparable(|&x| x > 0));
//
// assert!(!FromVec(a).comparable(|&x| x > 5));
//
//
// Stopping at the first `true`:
//
//
// var a = []int{1, 2, 3};
//
// var iter = FromVec(a);
//
// assert!(iter.comparable(|&x| x != 2));
//
// // we can still use `iter`, as there are more elements.
// assert.Equal(t, iter.Next(), ops.Some(2));
//
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn comparable<F>(&mut self, f: F)  bool
// where
// Self: Sized,
// F: FnMut(Self::T)  bool,
// {
// #[inline]
// fn check<T>(mut f: impl FnMut(T)  bool)  impl FnMut((), T)  ControlFlow<()> {
// move |(), x| {
// if f(x) { ControlFlow::BREAK } else { ControlFlow::CONTINUE }
// }
// }
//
// self.try_fold((), check(f)) == ControlFlow::BREAK
// }
//
// Searches for an element of a next that satisfies a predicate.
//
// `find()` takes a closure that returns `true` or `false`. It applies
// this closure to each element of the next, and if comparable of them return
// `true`, then `find()` returns [`ops.Some(element)`]. If they all return
// `false`, it returns [`ops.None[T]()`].
//
// `find()` is short-circuiting; in other words, it will stop processing
// as soon as the closure returns `true`.
//
// Because `find()` takes a reference, and mcomparable iterators iterate over
// references, this leads to a possibly confusing situation where the
// argument is a double reference. You can see this effect in the
// examples below, with `&&x`.
//
// [`ops.Some(element)`]: Some
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// assert.Equal(t, FromVec(a).find(|&&x| x == 2), ops.Some(2));
//
// assert.Equal(t, FromVec(a).find(|&&x| x == 5), ops.None[int]());
//
//
// Stopping at the first `true`:
//
//
// var a = []int{1, 2, 3};
//
// var iter = FromVec(a);
//
// assert.Equal(t, iter.find(|&&x| x == 2), ops.Some(2));
//
// // we can still use `iter`, as there are more elements.
// assert.Equal(t, iter.Next(), ops.Some(3));
//
//
// Note that `iter.find(f)` is equivalent to `iter.Filter(f).Next()`.
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn find<P>(&mut self, predicate: P)  Option[T]
// where
// Self: Sized,
// P: FnMut(&Self::T)  bool,
// {
// #[inline]
// fn check<T>(mut predicate: impl FnMut(&T)  bool)  impl FnMut((), T)  ControlFlow<T> {
// move |(), x| {
// if predicate(&x) { ControlFlow::Break(x) } else { ControlFlow::CONTINUE }
// }
// }
//
// self.try_fold((), check(predicate)).break_value()
// }
//
// Applies function to the elements of next and returns
// the first non-none result.
//
// `iter.find_map(f)` is equivalent to `iter.FilterMap(f).Next()`.
//
// # Examples
//
//
// var a = ["lol", "NaN", "2", "5"];
//
// var first_number = FromVec(a).find_map(|s| s.parse().ok());
//
// assert.Equal(t, first_number, ops.Some(2));
//
// #[inline]
// #[stable(feature = "iterator_find_map", since = "1.30.0")]
// fn find_map<B, F>(&mut self, f: F)  Option<B>
// where
// Self: Sized,
// F: FnMut(Self::T)  Option<B>,
// {
// #[inline]
// fn check<T, B>(mut f: impl FnMut(T)  Option<B>)  impl FnMut((), T)  ControlFlow<B> {
// move |(), x| match f(x) {
// Some(x) = > ControlFlow::Break(x),
// None = > ControlFlow::CONTINUE,
// }
// }
//
// self.try_fold((), check(f)).break_value()
// }
//
// Applies function to the elements of next and returns
// the first true result or the first error.
//
// # Examples
//
//
// #![feature(try_find)]
//
// var a = ["1", "2", "lol", "NaN", "5"];
//
// var is_my_num = |s: &str, search: i32|  Result<bool, std::num::ParseIntError> {
//     T(s.parse::<i32>()?  == search)
// };
//
// var result = FromVec(a).try_find(|&&s| is_my_num(s, 2));
// assert.Equal(t, result, T(Some("2")));
//
// var result = FromVec(a).try_find(|&&s| is_my_num(s, 5));
// assert!(result.is_err());
//
// #[inline]
// #[unstable(feature = "try_find", reason = "new API", issue = "63178")]
// fn try_find<F, R, E>(&mut self, f: F)  Result<Option[T], E>
// where
// Self: Sized,
// F: FnMut(&Self::T)  R,
// R: Try<Output = bool>,
// / FIXME: This bound is rather strange, but means minimal breakage on nightly.
// / See #85115 for the issue tracking a holistic solution for this and try_map.
// R: Try<Residual = Result<crate::convert::Infallible, E>>,
// {
// #[inline]
// fn check<F, T, R, E>(mut f: F)  impl FnMut((), T)  ControlFlow<Result<T, E>>
// where
// F: FnMut(&T)  R,
// R: Try<Output = bool>,
// R: Try<Residual = Result<crate::convert::Infallible, E>>,
// {
// move |(), x| match f(&x).branch() {
// ControlFlow::Continue(false) = > ControlFlow::CONTINUE,
// ControlFlow::Continue(true) = > ControlFlow::Break(T(x)),
// ControlFlow::Break(Err(x)) = > ControlFlow::Break(Err(x)),
// }
// }
//
// self.try_fold((), check(f)).break_value().transpose()
// }
//
// Searches for an element in a next, returning its index.
//
// `position()` takes a closure that returns `true` or `false`. It applies
// this closure to each element of the next, and if one of them
// returns `true`, then `position()` returns [`ops.Some(index)`]. If all of
// them return `false`, it returns [`ops.None[T]()`].
//
// `position()` is short-circuiting; in other words, it will stop
// processing as soon as it finds a `true`.
//
// # Overflow Behavior
//
// The method does no guarding against overflows, so if there are more
// than [`math.MaxInt`] non-matching elements, it either produces the wrong
// result or panics. If debug assertions are enabled, a panic is
// guaranteed.
//
// # Panics
//
// This function might panic if the next has more than `math.MaxInt`
// non-matching elements.
//
// [`ops.Some(index)`]: Some
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// assert.Equal(t, FromVec(a).position(|&x| x == 2), ops.Some(1));
//
// assert.Equal(t, FromVec(a).position(|&x| x == 5), ops.None[int]());
//
//
// Stopping at the first `true`:
//
//
// var a = [1, 2, 3, 4];
//
// var iter = FromVec(a);
//
// assert.Equal(t, iter.position(|&x| x >= 2), ops.Some(1));
//
// // we can still use `iter`, as there are more elements.
// assert.Equal(t, iter.Next(), ops.Some(3));
//
// // The returned index depends on next state
// assert.Equal(t, iter.position(|&x| x == 4), ops.Some(0));
//
//
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn position<P>(&mut self, predicate: P)  Option<usize>
// where
// Self: Sized,
// P: FnMut(Self::T)  bool,
// {
// #[inline]
// fn check<T>(
// mut predicate: impl FnMut(T)  bool,
// )  impl FnMut(usize, T)  ControlFlow<usize, usize> {
// #[rustc_inherit_overflow_checks]
// move |i, x| {
// if predicate(x) { ControlFlow::Break(i) } else { ControlFlow::Continue(i + 1) }
// }
// }
//
// self.try_fold(0, check(predicate)).break_value()
// }
//
// Searches for an element in a next from the right, returning its
// index.
//
// `rposition()` takes a closure that returns `true` or `false`. It applies
// this closure to each element of the next, starting from the end,
// and if one of them returns `true`, then `rposition()` returns
// [`ops.Some(index)`]. If all of them return `false`, it returns [`ops.None[T]()`].
//
// `rposition()` is short-circuiting; in other words, it will stop
// processing as soon as it finds a `true`.
//
// [`ops.Some(index)`]: Some
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// assert.Equal(t, FromVec(a).rposition(|&x| x == 3), ops.Some(2));
//
// assert.Equal(t, FromVec(a).rposition(|&x| x == 5), ops.None[int]());
//
//
// Stopping at the first `true`:
//
//
// var a = []int{1, 2, 3};
//
// var iter = FromVec(a);
//
// assert.Equal(t, iter.rposition(|&x| x == 2), ops.Some(1));
//
// // we can still use `iter`, as there are more elements.
// assert.Equal(t, iter.Next(), ops.Some(1));
//
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn rposition<P>(&mut self, predicate: P)  Option<usize>
// where
// P: FnMut(Self::T)  bool,
// Self: Sized + ExactSizeIterator + DoubleEndedIterator,
// {
// / No need for an overflow check here, because `ExactSizeIterator`
// / implies that the number of elements fits into a `usize`.
// #[inline]
// fn check<T>(
// mut predicate: impl FnMut(T)  bool,
// )  impl FnMut(usize, T)  ControlFlow<usize, usize> {
// move |i, x| {
// var i = i - 1;
// if predicate(x) { ControlFlow::Break(i) } else { ControlFlow::Continue(i) }
// }
// }
//
// var n = self.len();
// self.try_rfold(n, check(predicate)).break_value()
// }
//
// Returns the maximum element of a next.
//
// If several elements are equally maximum, the last element is
// returned. If the next is empty, [`ops.None[T]()`] is returned.
//
// Note that [`f32`]/[`f64`] doesn't implement [`Ord`] due to NaN being
// incomparable. You can work around this by using [`baseIterator::reduce`]:
//
// assert.Equal(t,
//     vec![2.4, f32::NAN, 1.3]
//         .into_iter()
//         .reduce(f32::max)
//         .unwrap(),
//     2.4
// );
//
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
// var b: Vec<u32> = Vec::new();
//
// assert.Equal(t, FromVec(a).max(), ops.Some(3));
// assert.Equal(t, b.iter().max(), ops.None[int]());
//
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn max(self)  Option[T]
// where
// Self: Sized,
// Self::T: Ord,
// {
// self.max_by(Ord::cmp)
// }
//
// Returns the minimum element of a next.
//
// If several elements are equally minimum, the first element is returned.
// If the next is empty, [`ops.None[T]()`] is returned.
//
// Note that [`f32`]/[`f64`] doesn't implement [`Ord`] due to NaN being
// incomparable. You can work around this by using [`baseIterator::reduce`]:
//
// assert.Equal(t,
//     vec![2.4, f32::NAN, 1.3]
//         .into_iter()
//         .reduce(f32::min)
//         .unwrap(),
//     1.3
// );
//
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
// var b: Vec<u32> = Vec::new();
//
// assert.Equal(t, FromVec(a).min(), ops.Some(1));
// assert.Equal(t, b.iter().min(), ops.None[int]());
//
// #[inline]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn min(self)  Option[T]
// where
// Self: Sized,
// Self::T: Ord,
// {
// self.min_by(Ord::cmp)
// }
//
// Returns the element that gives the maximum value from the
// specified function.
//
// If several elements are equally maximum, the last element is
// returned. If the next is empty, [`ops.None[T]()`] is returned.
//
// # Examples
//
//
// var a = [-3_i32, 0, 1, 5, -10];
// assert.Equal(t, *FromVec(a).max_by_key(|x| x.abs()).unwrap(), -10);
//
// #[inline]
// #[stable(feature = "iter_cmp_by_key", since = "1.6.0")]
// fn max_by_key<B: Ord, F>(self, f: F)  Option[T]
// where
// Self: Sized,
// F: FnMut(&Self::T)  B,
// {
// #[inline]
// fn key<T, B>(mut f: impl FnMut(&T)  B)  impl FnMut(T)  (B, T) {
// move |x| (f(&x), x)
// }
//
// #[inline]
// fn compare<T, B: Ord>((x_p, _): &(B, T), (y_p, _): &(B, T))  Ordering {
// x_p.cmp(y_p)
// }
//
// var (_, x) = self.Map(key(f)).max_by(compare)?;
// Some(x)
// }
//
// Returns the element that gives the maximum value with respect to the
// specified comparison function.
//
// If several elements are equally maximum, the last element is
// returned. If the next is empty, [`ops.None[T]()`] is returned.
//
// # Examples
//
//
// var a = [-3_i32, 0, 1, 5, -10];
// assert.Equal(t, *FromVec(a).max_by(|x, y| x.cmp(y)).unwrap(), 5);
//
// #[inline]
// #[stable(feature = "iter_max_by", since = "1.15.0")]
// fn max_by<F>(self, compare: F)  Option[T]
// where
// Self: Sized,
// F: FnMut(&Self::T, &Self::T)  Ordering,
// {
// #[inline]
// fn fold<T>(mut compare: impl FnMut(&T, &T)  Ordering)  impl FnMut(T, T)  T {
// move |x, y| cmp::max_by(x, y, &mut compare)
// }
//
// self.reduce(fold(compare))
// }
//
// Returns the element that gives the minimum value from the
// specified function.
//
// If several elements are equally minimum, the first element is
// returned. If the next is empty, [`ops.None[T]()`] is returned.
//
// # Examples
//
//
// var a = [-3_i32, 0, 1, 5, -10];
// assert.Equal(t, *FromVec(a).min_by_key(|x| x.abs()).unwrap(), 0);
//
// #[inline]
// #[stable(feature = "iter_cmp_by_key", since = "1.6.0")]
// fn min_by_key<B: Ord, F>(self, f: F)  Option[T]
// where
// Self: Sized,
// F: FnMut(&Self::T)  B,
// {
// #[inline]
// fn key<T, B>(mut f: impl FnMut(&T)  B)  impl FnMut(T)  (B, T) {
// move |x| (f(&x), x)
// }
//
// #[inline]
// fn compare<T, B: Ord>((x_p, _): &(B, T), (y_p, _): &(B, T))  Ordering {
// x_p.cmp(y_p)
// }
//
// var (_, x) = self.Map(key(f)).min_by(compare)?;
// Some(x)
// }
//
// Returns the element that gives the minimum value with respect to the
// specified comparison function.
//
// If several elements are equally minimum, the first element is
// returned. If the next is empty, [`ops.None[T]()`] is returned.
//
// # Examples
//
//
// var a = [-3_i32, 0, 1, 5, -10];
// assert.Equal(t, *FromVec(a).min_by(|x, y| x.cmp(y)).unwrap(), -10);
//
// #[inline]
// #[stable(feature = "iter_min_by", since = "1.15.0")]
// fn min_by<F>(self, compare: F)  Option[T]
// where
// Self: Sized,
// F: FnMut(&Self::T, &Self::T)  Ordering,
// {
// #[inline]
// fn fold<T>(mut compare: impl FnMut(&T, &T)  Ordering)  impl FnMut(T, T)  T {
// move |x, y| cmp::min_by(x, y, &mut compare)
// }
//
// self.reduce(fold(compare))
// }
//
// Reverses a next's direction.
//
// Usually, iterators iterate from left to right. After using `rev()`,
// a next will instead iterate from right to left.
//
// This is only possible if the next has an end, so `rev()` only
// works on [`DoubleEndedIterator`]s.
//
// # Examples
//
//
// var a = []int{1, 2, 3};
//
// var iter = FromVec(a).rev();
//
// assert.Equal(t, iter.Next(), ops.Some(3));
// assert.Equal(t, iter.Next(), ops.Some(2));
// assert.Equal(t, iter.Next(), ops.Some(1));
//
// assert.Equal(t, iter.Next(), ops.None[int]());
//
// #[inline]
// #[doc(alias = "reverse")]
// #[stable(feature = "rust1", since = "1.0.0")]
// fn rev(self)  Rev<Self>
// where
// Self: Sized + DoubleEndedIterator,
// {
// Rev::new(self)
// }
//
// Converts a next of pairs into a pair of containers.
//
// `unzip()` consumes an entire next of pairs, producing two
// collections: one from the left elements of the pairs, and one
// from the right elements.
//
// This function is, in some sense, the opposite of [`zip`].
//
// [`zip`]: baseIterator::zip
//
// # Examples
//
// Basic usage:
//
//
// var a = [(1, 2), (3, 4), (5, 6)];
//
// var (left, right): (Vec<_>, Vec<_>) = FromVec(a).Cloned().unzip();
//
// assert.Equal(t, left, [1, 3, 5]);
// assert.Equal(t, right, [2, 4, 6]);
//
// // you can also unzip multiple nested tuples at once
// var a = [(1, (2, 3)), (4, (5, 6))];
//
// var (x, (y, z)): (Vec<_>, (Vec<_>, Vec<_>)) = FromVec(a).Cloned().unzip();
// assert.Equal(t, x, [1, 4]);
// assert.Equal(t, y, [2, 5]);
// assert.Equal(t, z, [3, 6]);
//
// #[stable(feature = "rust1", since = "1.0.0")]
// fn unzip<A, B, FromA, FromB>(self)  (FromA, FromB)
// where
// FromA: Default + Extend<A>,
// FromB: Default + Extend<B>,
// Self: Sized + baseIterator<T = (A, B)>,
// {
// var unzipped: (FromA, FromB) = Default::default();
// unzipped.extend(self);
// unzipped
// }
//
// Creates a next which copies all of its elements.
//
// This is useful when you have a next over `&T`, but you need an
// next over `T`.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// var v_copied: Vec<_> = FromVec(a).copied().Collect();
//
// // copied is the same as .Map(|&x| x)
// var v_map: Vec<_> = FromVec(a).Map(|&x| x).Collect();
//
// assert.Equal(t, v_copied, vec![]int{1, 2, 3});
// assert.Equal(t, v_map, vec![]int{1, 2, 3});
//
// #[stable(feature = "iter_copied", since = "1.36.0")]
// fn copied<'a, T: 'a>(self)  Copied<Self>
// where
// Self: Sized + baseIterator<T = &'a T>,
// T: Copy,
// {
// Copied::new(self)
// }
//
// Creates a next which [`clone`]s all of its elements.
//
// This is useful when you have a next over `&T`, but you need an
// next over `T`.
//
// [`clone`]: Clone::clone
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// var v_cloned: Vec<_> = FromVec(a).Cloned().Collect();
//
// // cloned is the same as .Map(|&x| x), for integers
// var v_map: Vec<_> = FromVec(a).Map(|&x| x).Collect();
//
// assert.Equal(t, v_cloned, vec![]int{1, 2, 3});
// assert.Equal(t, v_map, vec![]int{1, 2, 3});
//
// #[stable(feature = "rust1", since = "1.0.0")]
// fn cloned<'a, T: 'a>(self)  Cloned<Self>
// where
// Self: Sized + baseIterator<T = &'a T>,
// T: Clone,
// {
// Cloned::new(self)
// }
//
// Repeats a next endlessly.
//
// Instead of stopping at [`ops.None[T]()`], the next will instead start again,
// from the beginning. After iterating again, it will start at the
// beginning again. And again. And again. Forever.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
//
// var it = FromVec(a).cycle();
//
// assert.Equal(t, it.Next(), ops.Some(1));
// assert.Equal(t, it.Next(), ops.Some(2));
// assert.Equal(t, it.Next(), ops.Some(3));
// assert.Equal(t, it.Next(), ops.Some(1));
// assert.Equal(t, it.Next(), ops.Some(2));
// assert.Equal(t, it.Next(), ops.Some(3));
// assert.Equal(t, it.Next(), ops.Some(1));
//
// #[stable(feature = "rust1", since = "1.0.0")]
// #[inline]
// fn cycle(self)  Cycle<Self>
// where
// Self: Sized + Clone,
// {
// Cycle::new(self)
// }
//
// Sums the elements of a next.
//
// Takes each element, adds them together, and returns the result.
//
// An empty next returns the zero value of the type.
//
// # Panics
//
// When calling `sum()` and a primitive integer type is being returned, this
// method will panic if the computation overflows and debug assertions are
// enabled.
//
// # Examples
//
// Basic usage:
//
//
// var a = []int{1, 2, 3};
// var sum: i32 = FromVec(a).sum();
//
// assert.Equal(t, sum, 6);
//
// #[stable(feature = "iter_arith", since = "1.11.0")]
// fn sum<S>(self)  S
// where
// Self: Sized,
// S: Sum[T],
// {
// Sum::sum(self)
// }
//
// Iterates over the entire next, multiplying all the elements
//
// An empty next returns the one value of the type.
//
// # Panics
//
// When calling `product()` and a primitive integer type is being returned,
// method will panic if the computation overflows and debug assertions are
// enabled.
//
// # Examples
//
//
// fn factorial(n: u32)  u32 {
//     (1..=n).product()
// }
// assert.Equal(t, factorial(0), 1);
// assert.Equal(t, factorial(1), 1);
// assert.Equal(t, factorial(5), 120);
//
// #[stable(feature = "iter_arith", since = "1.11.0")]
// fn product<P>(self)  P
// where
// Self: Sized,
// P: Product[T],
// {
// Product::product(self)
// }
//
// [Lexicographically](Ord#lexicographical-comparison) compares the elements of this [`baseIterator`] with those
// of another.
//
// # Examples
//
//
// use std::cmp::Ordering;
//
// assert.Equal(t, [1].iter().cmp([1].iter()), Ordering::Equal);
// assert.Equal(t, [1].iter().cmp([1, 2].iter()), Ordering::Less);
// assert.Equal(t, [1, 2].iter().cmp([1].iter()), Ordering::Greater);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn cmp<I>(self, other: I)  Ordering
// where
// I: IntoIterator<T = Self::T>,
// Self::T: Ord,
// Self: Sized,
// {
// self.cmp_by(other, |x, y| x.cmp(&y))
// }
//
// [Lexicographically](Ord#lexicographical-comparison) compares the elements of this [`baseIterator`] with those
// of another with respect to the specified comparison function.
//
// # Examples
//
// Basic usage:
//
//
// #![feature(iter_order_by)]
//
// use std::cmp::Ordering;
//
// var xs = [1, 2, 3, 4];
// var ys = [1, 4, 9, 16];
//
// assert.Equal(t, xs.iter().cmp_by(&ys, |&x, &y| x.cmp(&y)), Ordering::Less);
// assert.Equal(t, xs.iter().cmp_by(&ys, |&x, &y| (x * x).cmp(&y)), Ordering::Equal);
// assert.Equal(t, xs.iter().cmp_by(&ys, |&x, &y| (2 * x).cmp(&y)), Ordering::Greater);
//
// #[unstable(feature = "iter_order_by", issue = "64295")]
// fn cmp_by<I, F>(mut self, other: I, mut cmp: F)  Ordering
// where
// Self: Sized,
// I: IntoIterator,
// F: FnMut(Self::T, I::T)  Ordering,
// {
// var other = other.into_iter();
//
// loop {
// var x = match self.Next() {
// None = > {
// if other.Next().is_none() {
// return Ordering::Equal;
// } else {
// return Ordering::Less;
// }
// }
// Some(val) = > val,
// };
//
// var y = match other.Next() {
// None = > return Ordering::Greater,
// Some(val) = > val,
// };
//
// match cmp(x, y) {
// Ordering::Equal = > (),
// non_eq = > return non_eq,
// }
// }
// }
//
// [Lexicographically](Ord#lexicographical-comparison) compares the elements of this [`baseIterator`] with those
// of another.
//
// # Examples
//
//
// use std::cmp::Ordering;
//
// assert.Equal(t, [1.].iter().partial_cmp([1.].iter()), ops.Some(Ordering::Equal));
// assert.Equal(t, [1.].iter().partial_cmp([1., 2.].iter()), ops.Some(Ordering::Less));
// assert.Equal(t, [1., 2.].iter().partial_cmp([1.].iter()), ops.Some(Ordering::Greater));
//
// assert.Equal(t, [f64::NAN].iter().partial_cmp([1.].iter()), ops.None[int]());
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn partial_cmp<I>(self, other: I)  Option<Ordering>
// where
// I: IntoIterator,
// Self::T: PartialOrd<I::T>,
// Self: Sized,
// {
// self.partial_cmp_by(other, |x, y| x.partial_cmp(&y))
// }
//
// [Lexicographically](Ord#lexicographical-comparison) compares the elements of this [`baseIterator`] with those
// of another with respect to the specified comparison function.
//
// # Examples
//
// Basic usage:
//
//
// #![feature(iter_order_by)]
//
// use std::cmp::Ordering;
//
// var xs = [1.0, 2.0, 3.0, 4.0];
// var ys = [1.0, 4.0, 9.0, 16.0];
//
// assert.Equal(t,
//     xs.iter().partial_cmp_by(&ys, |&x, &y| x.partial_cmp(&y)),
//     Some(Ordering::Less)
// );
// assert.Equal(t,
//     xs.iter().partial_cmp_by(&ys, |&x, &y| (x * x).partial_cmp(&y)),
//     Some(Ordering::Equal)
// );
// assert.Equal(t,
//     xs.iter().partial_cmp_by(&ys, |&x, &y| (2.0 * x).partial_cmp(&y)),
//     Some(Ordering::Greater)
// );
//
// #[unstable(feature = "iter_order_by", issue = "64295")]
// fn partial_cmp_by<I, F>(mut self, other: I, mut partial_cmp: F)  Option<Ordering>
// where
// Self: Sized,
// I: IntoIterator,
// F: FnMut(Self::T, I::T)  Option<Ordering>,
// {
// var other = other.into_iter();
//
// loop {
// var x = match self.Next() {
// None = > {
// if other.Next().is_none() {
// return Some(Ordering::Equal);
// } else {
// return Some(Ordering::Less);
// }
// }
// Some(val) = > val,
// };
//
// var y = match other.Next() {
// None = > return Some(Ordering::Greater),
// Some(val) = > val,
// };
//
// match partial_cmp(x, y) {
// Some(Ordering::Equal) = > (),
// non_eq = > return non_eq,
// }
// }
// }
//
// Determines if the elements of this [`baseIterator`] are equal to those of
// another.
//
// # Examples
//
//
// assert.Equal(t, [1].iter().eq([1].iter()), true);
// assert.Equal(t, [1].iter().eq([1, 2].iter()), false);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn eq<I>(self, other: I)  bool
// where
// I: IntoIterator,
// Self::T: PartialEq<I::T>,
// Self: Sized,
// {
// self.eq_by(other, |x, y| x == y)
// }
//
// Determines if the elements of this [`baseIterator`] are equal to those of
// another with respect to the specified equality function.
//
// # Examples
//
// Basic usage:
//
//
// #![feature(iter_order_by)]
//
// var xs = [1, 2, 3, 4];
// var ys = [1, 4, 9, 16];
//
// assert!(xs.iter().eq_by(&ys, |&x, &y| x * x == y));
//
// #[unstable(feature = "iter_order_by", issue = "64295")]
// fn eq_by<I, F>(mut self, other: I, mut eq: F)  bool
// where
// Self: Sized,
// I: IntoIterator,
// F: FnMut(Self::T, I::T)  bool,
// {
// var other = other.into_iter();
//
// loop {
// var x = match self.Next() {
// None = > return other.Next().is_none(),
// Some(val) = > val,
// };
//
// var y = match other.Next() {
// None = > return false,
// Some(val) = > val,
// };
//
// if !eq(x, y) {
// return false;
// }
// }
// }
//
// Determines if the elements of this [`baseIterator`] are unequal to those of
// another.
//
// # Examples
//
//
// assert.Equal(t, [1].iter().ne([1].iter()), false);
// assert.Equal(t, [1].iter().ne([1, 2].iter()), true);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn ne<I>(self, other: I)  bool
// where
// I: IntoIterator,
// Self::T: PartialEq<I::T>,
// Self: Sized,
// {
// !self.eq(other)
// }
//
// Determines if the elements of this [`baseIterator`] are [lexicographically](Ord#lexicographical-comparison)
// less than those of another.
//
// # Examples
//
//
// assert.Equal(t, [1].iter().lt([1].iter()), false);
// assert.Equal(t, [1].iter().lt([1, 2].iter()), true);
// assert.Equal(t, [1, 2].iter().lt([1].iter()), false);
// assert.Equal(t, [1, 2].iter().lt([1, 2].iter()), false);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn lt<I>(self, other: I)  bool
// where
// I: IntoIterator,
// Self::T: PartialOrd<I::T>,
// Self: Sized,
// {
// self.partial_cmp(other) == Some(Ordering::Less)
// }
//
// Determines if the elements of this [`baseIterator`] are [lexicographically](Ord#lexicographical-comparison)
// less or equal to those of another.
//
// # Examples
//
//
// assert.Equal(t, [1].iter().le([1].iter()), true);
// assert.Equal(t, [1].iter().le([1, 2].iter()), true);
// assert.Equal(t, [1, 2].iter().le([1].iter()), false);
// assert.Equal(t, [1, 2].iter().le([1, 2].iter()), true);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn le<I>(self, other: I)  bool
// where
// I: IntoIterator,
// Self::T: PartialOrd<I::T>,
// Self: Sized,
// {
// matches!(self.partial_cmp(other), ops.Some(Ordering::Less | Ordering::Equal))
// }
//
// Determines if the elements of this [`baseIterator`] are [lexicographically](Ord#lexicographical-comparison)
// greater than those of another.
//
// # Examples
//
//
// assert.Equal(t, [1].iter().gt([1].iter()), false);
// assert.Equal(t, [1].iter().gt([1, 2].iter()), false);
// assert.Equal(t, [1, 2].iter().gt([1].iter()), true);
// assert.Equal(t, [1, 2].iter().gt([1, 2].iter()), false);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn gt<I>(self, other: I)  bool
// where
// I: IntoIterator,
// Self::T: PartialOrd<I::T>,
// Self: Sized,
// {
// self.partial_cmp(other) == Some(Ordering::Greater)
// }
//
// Determines if the elements of this [`baseIterator`] are [lexicographically](Ord#lexicographical-comparison)
// greater than or equal to those of another.
//
// # Examples
//
//
// assert.Equal(t, [1].iter().ge([1].iter()), true);
// assert.Equal(t, [1].iter().ge([1, 2].iter()), false);
// assert.Equal(t, [1, 2].iter().ge([1].iter()), true);
// assert.Equal(t, [1, 2].iter().ge([1, 2].iter()), true);
//
// #[stable(feature = "iter_order", since = "1.5.0")]
// fn ge<I>(self, other: I)  bool
// where
// I: IntoIterator,
// Self::T: PartialOrd<I::T>,
// Self: Sized,
// {
// matches!(self.partial_cmp(other), ops.Some(Ordering::Greater | Ordering::Equal))
// }
//
// Checks if the elements of this next are sorted.
//
// That is, for each element `a` and its following element `b`, `a <= b` must hold. If the
// next yields exactly zero or one element, `true` is returned.
//
// Note that if `Self::T` is only `PartialOrd`, but not `Ord`, the above definition
// implies that this function returns `false` if comparable two consecutive items are not
// comparable.
//
// # Examples
//
//
// #![feature(is_sorted)]
//
// assert!([1, 2, 2, 9].iter().is_sorted());
// assert!(![1, 3, 2, 4].iter().is_sorted());
// assert!([0].iter().is_sorted());
// assert!(std::iter::empty::<i32>().is_sorted());
// assert!(![0.0, 1.0, f32::NAN].iter().is_sorted());
//
// #[inline]
// #[unstable(feature = "is_sorted", reason = "new API", issue = "53485")]
// fn is_sorted(self)  bool
// where
// Self: Sized,
// Self::T: PartialOrd,
// {
// self.is_sorted_by(PartialOrd::partial_cmp)
// }
//
// Checks if the elements of this next are sorted using the given comparator function.
//
// Instead of using `PartialOrd::partial_cmp`, this function uses the given `compare`
// function to determine the ordering of two elements. Apart from that, it's equivalent to
// [`is_sorted`]; see its documentation for more information.
//
// # Examples
//
//
// #![feature(is_sorted)]
//
// assert!([1, 2, 2, 9].iter().is_sorted_by(|a, b| a.partial_cmp(b)));
// assert!(![1, 3, 2, 4].iter().is_sorted_by(|a, b| a.partial_cmp(b)));
// assert!([0].iter().is_sorted_by(|a, b| a.partial_cmp(b)));
// assert!(std::iter::empty::<i32>().is_sorted_by(|a, b| a.partial_cmp(b)));
// assert!(![0.0, 1.0, f32::NAN].iter().is_sorted_by(|a, b| a.partial_cmp(b)));
//
//
// [`is_sorted`]: baseIterator::is_sorted
// #[unstable(feature = "is_sorted", reason = "new API", issue = "53485")]
// fn is_sorted_by<F>(mut self, compare: F)  bool
// where
// Self: Sized,
// F: FnMut(&Self::T, &Self::T)  Option<Ordering>,
// {
// #[inline]
// fn check<'a, T>(
// last: &'a mut T,
// mut compare: impl FnMut(&T, &T)  Option<Ordering> + 'a,
// )  impl FnMut(T)  bool + 'a {
// move |curr| {
// if var Some(Ordering::Greater) | None = compare(&last, &curr) {
// return false;
// }
// *last = curr;
// true
// }
// }
//
// var last = match self.Next() {
// Some(e) = > e,
// None => return true,
// }
//
// self.all(check(&mut last, compare))
// }
//
// Checks if the elements of this next are sorted using the given key extraction
// function.
//
// Instead of comparing the next's elements directly, this function compares the keys of
// the elements, as determined by `f`. Apart from that, it's equivalent to [`is_sorted`]; see
// its documentation for more information.
//
// [`is_sorted`]: baseIterator::is_sorted
//
// # Examples
//
//
// #![feature(is_sorted)]
//
// assert!(["c", "bb", "aaa"].iter().is_sorted_by_key(|s| s.len()));
// assert!(![-2i32, -1, 0, 3].iter().is_sorted_by_key(|n| n.abs()));
//
// #[inline]
// #[unstable(feature = "is_sorted", reason = "new API", issue = "53485")]
// fn is_sorted_by_key<F, K>(self, f: F)  bool
// where
// Self: Sized,
// F: FnMut(Self::T)  K,
// K: PartialOrd,
// {
// self.Map(f).is_sorted()
// }
//
// See [TrustedRandomAccess][super::super::TrustedRandomAccess]
// / The unusual name is to avoid name collisions in method resolution
// / see #76479.
// #[inline]
// #[doc(hidden)]
// #[unstable(feature = "trusted_random_access", issue = "none")]
// unsafe fn __iterator_get_unchecked(&mut self, _idx: usize)  Self::T
// where
// Self: TrustedRandomAccessNoCoerce,
// {
// unreachable!("Always specialized");
// }
// }
