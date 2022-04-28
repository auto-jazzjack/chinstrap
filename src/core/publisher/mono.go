package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
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

func (m Mono[T]) Map(consumer func(T) T) Mono[T] {
	return NewMonoMap(m, consumer)
}

func (m Mono[T]) Subscribe() {
	m.actual.Subscribe(NewLamdaSubscriber[T](nil, nil, nil))
}

func (m Mono[T]) SubscribeCore(subscriber core.CoreSubscriber[T]) {
	panic("shold be implmented")
}

func Subscribe0[T any](m core.CorePublisher[T], s reactive.Subscriber[T]) {
	pub := core.CorePublisher[T](m)
	sub := s.(core.CoreSubscriber[T])
	pub.SubscribeCore(sub)
}
