package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoJust[T any] struct {
	value T
}

type MonoJustImpl[T any] interface {
	Fusable
	Mono[T]
}

func NewMonoJust[T any](t T) *Wrapper[T] {
	v := &MonoJust[T]{
		value: t,
	}
	return NewWrapper[T](v)
}

func (m *MonoJust[T]) GetWrapper() *Wrapper[T] {
	//return parent? actually it is just wrapper
	return nil
}

func (m *MonoJust[T]) SubscribeCore(actual core.CoreSubscriber[T]) {
	actual.OnSubscribe(NewScalarSubscription(actual, m.value))
}

func (m *MonoJust[T]) Subscribe(s reactive.Subscriber[T]) {

}
