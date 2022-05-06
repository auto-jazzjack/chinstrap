package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
	"chinstrap/core/util"
)

type MonoMap[I util.All, O util.All] struct {
	mapper func(I) O
	mono   Mono[I]
}

func NewMonoMap[I util.All, O util.All](source Mono[I], mapper func(I) O) Mono[O] {
	v := &MonoMap[I, O]{
		mono:   source,
		mapper: mapper,
	}
	return Mono[O]{
		actual: v,
	}
}

func (m MonoMap[I, O]) SubscribeCore(actual core.CoreSubscriber[O]) {
	m.mono.actual.Subscribe(newMonoMapSubscriber(actual, m.mapper))
}

func (m MonoMap[I, O]) Subscribe(s reactive.Subscriber[O]) {
	Subscribe0(core.CorePublisher[O](m), s)
}

type MonoMapSubscriber[I util.All, O util.All] struct {
	mapper func(I) O
	src    core.CoreSubscriber[O]
	sub    reactive.Subscription
}

func newMonoMapSubscriber[I util.All, O util.All](m core.CoreSubscriber[O], mapper func(I) O) reactive.Subscriber[I] {
	return &MonoMapSubscriber[I, O]{
		mapper: mapper,
		src:    m,
	}
}

func (mm *MonoMapSubscriber[I, O]) OnSubscribe(s reactive.Subscription) {
	mm.sub = s
	mm.src.OnSubscribe(mm)
}
func (mm *MonoMapSubscriber[I, O]) OnError(t error) {
	mm.src.OnError(t)
}
func (mm *MonoMapSubscriber[I, O]) OnNext(t I) error {
	res := mm.mapper(t)
	return mm.src.OnNext(res)
}

func (mm *MonoMapSubscriber[I, O]) OnComplete() {
	mm.src.OnComplete()
}

func (mm *MonoMapSubscriber[I, O]) Request(n int64) {
	mm.sub.Request(n)
}
func (mm *MonoMapSubscriber[I, O]) Cancel() {
	mm.sub.Cancel()
}

func (mm *MonoMapSubscriber[I, O]) CurrentContext() {

}
