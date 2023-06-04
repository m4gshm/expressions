package opt

type Optional[T any] struct {
	element T
	ok      bool
}

func (o Optional[T]) GetOr(element T) T {
	if o.ok {
		return o.element
	}
	return element
}

func (o Optional[T]) Get() (T, bool) {
	return o.element, o.ok
}
