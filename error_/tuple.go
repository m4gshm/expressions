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

func (o Tuple[T]) Run(routine func(T)) Tuple[T] {
	if o.Ok() {
		routine(o.val)
	}
	return o
}

func (o Tuple[T]) RunErr(routine func(T) error) Tuple[T] {
	if o.Ok() {
		if err := routine(o.val); err != nil {
			return Tuple[T]{o.val, err}
		}
	}
	return o
}
