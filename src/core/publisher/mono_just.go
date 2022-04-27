package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoJustImpl[T any] struct {
	value T
}

/* type MonoJust[T any] interface {
	reactive.Subscriber[T]
}
*/
func NewMonoJust[T any](t T) Mono[T] {
	v := &MonoJustImpl[T]{
		value: t,
	}
	return Mono[T]{
		actual: v,
	}
}

func (m *MonoJustImpl[T]) SubscribeCore(actual core.CoreSubscriber[T]) {
	actual.OnSubscribe(NewScalarSubscription(actual, m.value))
}

func (m *MonoJustImpl[T]) Subscribe(s reactive.Subscriber[T]) {
	pub := core.CorePublisher[T](m)
	sub := s.(core.CoreSubscriber[T])
	pub.Subscribe(sub)
}
