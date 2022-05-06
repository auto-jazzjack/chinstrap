package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
	"chinstrap/core/util"
)

type InnerPublisher[T util.All] struct {
	m Mono0[T]
}

func (publisher InnerPublisher[T]) SubscribeCore(subscriber core.CoreSubscriber[util.All]) {
	publisher.m.SubscribeCore(subscriber.(core.CoreSubscriber[T]))
}

func (publisher InnerPublisher[T]) Subscribe(s reactive.Subscriber[util.All]) {
	publisher.m.Subscribe(s.(reactive.Subscriber[T]))
	//publisher.m.Subscribe(s)
}
func ConvertToObject[T util.All](m Mono[T]) core.CorePublisher[util.All] {
	return &InnerPublisher[T]{
		m: m,
	}
}
