package error_

type Tuple2[F, S any] struct {
	first  F
	second S
	err    error
}

func (o Tuple2[F, S]) Get() (F, S, error) {
	return o.first, o.second, o.err
}

func (o Tuple2[F, S]) Vals() (F, S) {
	return o.first, o.second
}

func (o Tuple2[F, S]) Err() error {
	return o.err
}

func (o Tuple2[F, S]) Ok() bool {
	return o.err != nil
}

func (o Tuple2[F, S]) Run(routine func(F, S)) Tuple2[F, S] {
	if o.Ok() {
		routine(o.first, o.second)
	}
	return o
}
