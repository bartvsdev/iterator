package iterator

import (
	"errors"
)

// Iterator represents iterator yielding elements of type T.
type Iterator[T any] interface {
	// Next yields a new value from the Iterator.
	Next() (T, error)
}

var ErrNoNext = errors.New("invoking Next on an empty iterator")

type sliceIter[T any] struct {
	slice []T
}

func (iter *sliceIter[T]) Next() (value T, err error) {
	if len(iter.slice) > 0 {
		value = iter.slice[0]
		iter.slice = iter.slice[1:]
	} else {
		err = ErrNoNext
	}
	return
}

// Slice returns an Iterator that yields elements from a slice.
func Slice[T any](slice []T) Iterator[T] {
	return &sliceIter[T]{
		slice: slice,
	}
}

type mapIter[T1, T2 any] struct {
	inner Iterator[T1]
	fn    func(T1) T2
}

func (iter *mapIter[T1, T2]) Next() (value T2, err error) {
	var oldValue T1
	if oldValue, err = iter.inner.Next(); err == nil {
		value = iter.fn(oldValue)
	}
	return
}

// Map returns an iterator of type T2 by applying fn to elements of iter.
func Map[T1, T2 any](iter Iterator[T1], fn func(T1) T2) Iterator[T2] {
	return &mapIter[T1, T2]{
		inner: iter,
		fn:    fn,
	}
}

type mapErrIter[T1, T2 any] struct {
	inner Iterator[T1]
	fn    func(T1) (T2, error)
}

func (iter *mapErrIter[T1, T2]) Next() (value T2, err error) {
	var oldValue T1
	if oldValue, err = iter.inner.Next(); err == nil {
		value, err = iter.fn(oldValue)
	}
	return
}

// MapErr returns an iterator of type T2 by applying a possibly erroneous fn to
// elements of iter.
func MapErr[T1, T2 any](iter Iterator[T1], fn func(T1) (T2, error)) Iterator[T2] {
	return &mapErrIter[T1, T2]{
		inner: iter,
		fn:    fn,
	}
}

type filterIter[T any] struct {
	inner Iterator[T]
	pred  func(T) bool
}

func (iter *filterIter[T]) Next() (value T, err error) {
	for value, err = iter.inner.Next(); err == nil; value, err = iter.inner.Next() {
		if iter.pred(value) {
			break
		}
	}
	return
}

// Filter returns an iterator of only the items in iter conform given predicate
func Filter[T any](iter Iterator[T], pred func(T) bool) Iterator[T] {
	return &filterIter[T]{
		inner: iter,
		pred:  pred,
	}
}

// ToSlice consumes the iterator yielding a slice from its values
func ToSlice[T any](iter Iterator[T]) ([]T, error) {
	result := []T{}
	value, err := iter.Next()
	for ; err == nil; value, err = iter.Next() {
		result = append(result, value)
	}
	if errors.Is(err, ErrNoNext) {
		return result, nil
	}
	return nil, err
}

// ForEach consumes the Iterator applying fn to each yielded value.
// Returns the first error encountered while iterating, or nil
func ForEach[T any](iter Iterator[T], fn func(T)) (err error) {
	var value T
	for value, err = iter.Next(); err == nil; value, err = iter.Next() {
		fn(value)
	}
	if errors.Is(err, ErrNoNext) {
		err = nil
	}
	return
}

// Fold consumes the iterator and returns a single value by incrementally
// applying the combining function to an accumulator and the next element
// of iter, starting with the initial value as the accumulator.
func Fold[T1 any, T2 any](iter Iterator[T1], init T2, comb func(T2, T1) T2) (T2, error) {
	acc := init
	value, err := iter.Next()
	for ; err == nil; value, err = iter.Next() {
		acc = comb(acc, value)
	}
	if errors.Is(err, ErrNoNext) {
		return acc, nil
	}
	return init, err
}
