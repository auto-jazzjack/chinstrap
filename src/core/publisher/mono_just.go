package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoJust[T any] struct {
	Mono  Mono[T]
	value T
}

type MonoJustImpl[T any] interface {
	Fusable
	core.CorePublisher[T]
	//SubscribeCore(subscriber core.CoreSubscriber[T])
	//Subscribe(s reactive.Subscriber[T])
}

func NewMonoJust[T any](t T) *MonoJust[T] {
	return &MonoJust[T]{
		value: t,
	}
}

func (m *MonoJust[T]) SubscribeCore(actual core.CoreSubscriber[T]) {
	actual.OnSubscribe(NewScalarSubscription(actual, m.value))
}

func (m *MonoJust[T]) Subscribe(s reactive.Subscriber[T]) {

}
