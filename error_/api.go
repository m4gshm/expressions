package error_

import (
	"errors"
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
