package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoFilter[I any] struct {
	predicate func(I) bool
	mono      Mono[I]
}

func NewMonoFilter[I any](source Mono[I], predicate func(I) bool) Mono[I] {
	v := &MonoFilter[I]{
		mono:      source,
		predicate: predicate,
	}
	return Mono[I]{
		actual: v,
	}
}

func (m *MonoFilter[I]) SubscribeCore(actual core.CoreSubscriber[I]) {
	m.mono.actual.Subscribe(newMonoFilterSubscriber(actual, m.predicate))
}

func (m *MonoFilter[I]) Subscribe(s reactive.Subscriber[I]) {
	Subscribe0(core.CorePublisher[I](m), s)
}

type MonoFilterSubscriber[I any] struct {
	predicate func(I) bool
	src       core.CoreSubscriber[I]
	sub       reactive.Subscription
}

func newMonoFilterSubscriber[I any](m core.CoreSubscriber[I], predicate func(I) bool) reactive.Subscriber[I] {
	return &MonoFilterSubscriber[I]{
		predicate: predicate,
		src:       m,
	}
}

func (mm *MonoFilterSubscriber[I]) OnSubscribe(s reactive.Subscription) {
	mm.sub = s
	mm.src.OnSubscribe(mm)
}
func (mm *MonoFilterSubscriber[I]) OnError(t error) {
	mm.src.OnError(t)
}
func (mm *MonoFilterSubscriber[I]) OnNext(t I) error {
	res := mm.predicate(t)
	if res {
		return mm.src.OnNext(t)
	}
	return nil
}

func (mm *MonoFilterSubscriber[I]) OnComplete() {
	mm.src.OnComplete()
}

func (mm *MonoFilterSubscriber[I]) Request(n int64) {
	mm.sub.Request(n)
}
func (mm *MonoFilterSubscriber[I]) Cancel() {
	mm.sub.Cancel()
}

func (mm *MonoFilterSubscriber[I]) CurrentContext() {

}
