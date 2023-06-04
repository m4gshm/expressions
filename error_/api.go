package error_

import "errors"


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
