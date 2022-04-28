package publisher

import (
	"chinstrap/core"
)

type Mono[T any] struct {
	actual Mono0[T]
	//Map() Mono[any]
}

type Mono0[T any] interface {
	core.CorePublisher[T]
}

func Just[V any](v V) Mono[V] {
	return NewMonoJust(v)

}

func (m Mono[T]) Map(consumer func(T, T) T) Mono[T] {
	return Mono[T]{}
}

func (m Mono[T]) Subscribe() {
	m.actual.SubscribeCore(NewLamdaSubscriber[T](nil, nil, nil))
}

func (m Mono[T]) SubscribeCore(subscriber core.CoreSubscriber[T]) {

}
