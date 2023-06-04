package opt

func Of[T any](element T, ok bool) Optional[T] {
	return Optional[T]{element: element, ok: ok}
}

func MapVal[M ~map[K]V, K comparable, V any](m M, key K) Optional[V] {
	val, ok := m[key]
	return Of(val, ok)
}
