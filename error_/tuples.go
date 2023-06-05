package error_

type Tuple[T any] struct {
	val T
	err error
}

func (o Tuple[T]) Get() (T, error) {
	return o.val, o.err
}

func (o Tuple[T]) Val() T {
	return o.val
}

func (o Tuple[T]) Err() error {
	return o.err
}

func (o Tuple[T]) Ok() bool {
	return o.err != nil
}
