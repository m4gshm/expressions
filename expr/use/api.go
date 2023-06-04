// Package use provides conditional expression builders
package use

import "github.com/m4gshm/expressions/opt"

// If builds use.If(condition, tru).Else(fals) expression builder.
// Looks like val := use.If(condition, valOnTrue).Else(defaltVal) tha can be rewrited by:
//
//	var val type
//	if condtion {
//		val = valOnTrue
//	} else {
//		val = defaltVal
//	}
func If[T any](condition bool, tru T) When[T] {
	return newWhen(condition, tru)
}

// Opt wraps optional result of a method/function
func Opt[T any](optional opt.Optional[T]) When[T] {
	element, ok := optional.Get()
	return newWhen(ok, element)
}

// MapVal retrieves the map value for the specified key, 'key exists' marker and wraps them into a condition expression object
func MapVal[M ~map[K]V, K comparable, V any](m M, key K) When[V] {
	v, ok := m[key]
	return newWhen(ok, v)
}

// IfGet is like If but aimed to use an getter function
func IfGet[T any](condition bool, then func() T) When[T] {
	return If(condition, evaluate(condition, then))
}

// IfGetErr is like If but aimed to use an error return function
func IfGetErr[T any](condition bool, tru func() (T, error)) WhenErr[T] {
	then, err := evaluateErr(condition, tru)
	return newWhenErr(condition, then, err)
}

// If_ is alias of IfErr
func If_[T any](condition bool, tru func() (T, error)) WhenErr[T] {
	return IfGetErr(condition, tru)
}
