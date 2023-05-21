package catch

import (
	errors "github.com/m4gshm/expressions/error_"
)

func Two[F, S any](first F, second S, err error) (*errors.Catcher, F, S) {
	return &errors.Catcher{Err: err}, first, second
}

func One[T any](result T, err error) (*errors.Catcher, T) {
	return &errors.Catcher{Err: err}, result
}
