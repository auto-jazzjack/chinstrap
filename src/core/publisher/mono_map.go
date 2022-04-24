package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoMap[I any, O any] struct {
	mapper func(I) O
	mono   Mono[I]
}

type MonoMapImpl[I any, O any] interface {
	core.CorePublisher[O]
}

func NewMonoMap[I any, O any](source Mono[I], mapper func(I) O) *Wrapper[O] {
	v := &MonoMap[I, O]{
		mono:   source,
		mapper: mapper,
	}
	return NewWrapper[O](v)
}

func (m *MonoMap[I, O]) SubscribeCore(actual core.CoreSubscriber[O]) {

}

func (m *MonoMap[I, O]) Subscribe(s reactive.Subscriber[O]) {

}
