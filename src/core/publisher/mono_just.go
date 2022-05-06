package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
	"chinstrap/core/util"
)

type MonoJustImpl[T util.All] struct {
	value T
}

/* type MonoJust[T any] interface {
	reactive.Subscriber[T]
}
*/
func NewMonoJust[T util.All](t T) Mono[T] {
	v := MonoJustImpl[T]{
		value: t,
	}

	return Mono[T]{
		actual: v,
	}
}

func (m MonoJustImpl[T]) SubscribeCore(actual core.CoreSubscriber[T]) {
	actual.OnSubscribe(NewScalarSubscription(actual, m.value))
}

func (m MonoJustImpl[T]) Subscribe(s reactive.Subscriber[T]) {
	Subscribe0(core.CorePublisher[T](m), s)
}
