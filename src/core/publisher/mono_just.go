package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoJust[T any] struct {
	value T
}

type MonoJustImpl[T any] interface {
	core.CorePublisher[T]
}

func NewMonoJust[T any](t T) Mono[T] {
	v := &MonoJust[T]{
		value: t,
	}
	return Mono[T]{
		actual: v,
	}
}

func (m *MonoJust[T]) SubscribeCore(actual core.CoreSubscriber[T]) {
	actual.OnSubscribe(NewScalarSubscription(actual, m.value))
}

func (m *MonoJust[T]) Subscribe(s reactive.Subscriber[T]) {

}
