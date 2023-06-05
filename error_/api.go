package error_

import (
	"errors"

	"github.com/m4gshm/expressions/expr/use"
)

func New(text string) error {
	return errors.New(text)
}

func As[T error](err error) (out T, ok bool) {
	return out, errors.As(err, &out)
}

func Check[T error](err error, predicate func(T) bool) bool {
	if out, ok := As[T](err); ok {
		return predicate(out)
	}
	return false
}

func Catch[T any](element T, err error) Tuple[T] {
	return Tuple[T]{element, err}
}

func Convert[From, To any](optional Tuple[From], converter func(From) To) Tuple[To] {
	var to To
	err := optional.err
	if err == nil {
		to = converter(optional.val)
	}
	return Tuple[To]{to, err}
}

func Convertt[From, To any](optional Tuple[From], converter func(From) (To, error)) Tuple[To] {
	var to To
	err := optional.err
	if err == nil {
		to, err = converter(optional.val)
	}
	return Tuple[To]{to, err}
}

func Zip[F, S any](first Tuple[F], second Tuple[S]) Tuple2[F, S] {
	var f F
	var s S
	err1 := first.err
	err2 := second.err
	err := use.If(err1 != nil, err1).Else(err2)
	if err1 == nil {
		f = first.val
	}
	if err2 == nil {
		s = second.val
	}
	return Tuple2[F, S]{f, s, err}
}
