package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoZip[I any, O any] struct {
	zipper func(...I) O
	monos  []Mono[I]
}

func NewMonoZip[I any, O any](sources []Mono[I], zipper func(...I) O) Mono[O] {
	v := &MonoZip[I, O]{
		monos:  sources,
		zipper: zipper,
	}
	return Mono[O]{
		actual: v,
	}
}

func (m *MonoZip[I, O]) SubscribeCore(actual core.CoreSubscriber[O]) {
	m.mono.actual.Subscribe(newMonoZipSubscriber(actual, m.Zipper))
}

func (m *MonoZip[I, O]) Subscribe(s reactive.Subscriber[O]) {
	Subscribe0(core.CorePublisher[O](m), s)
}

type MonoZipSubscriber[I any, O any] struct {
	Zipper func(I) O
	src    core.CoreSubscriber[O] //parent
	sub    reactive.Subscription
}

func newMonoZipSubscriber[I any, O any](m core.CoreSubscriber[O], Zipper func(I) O) reactive.Subscriber[I] {
	return &MonoZipSubscriber[I, O]{
		Zipper: Zipper,
		src:    m,
	}
}

func (mm *MonoZipSubscriber[I, O]) OnSubscribe(s reactive.Subscription) {
	mm.sub = s
	mm.src.OnSubscribe(mm)
}
func (mm *MonoZipSubscriber[I, O]) OnError(t error) {
	mm.src.OnError(t)
}
func (mm *MonoZipSubscriber[I, O]) OnNext(t I) error {
	res := mm.Zipper(t)
	return mm.src.OnNext(res)
}

func (mm *MonoZipSubscriber[I, O]) OnComplete() {
	mm.src.OnComplete()
}

func (mm *MonoZipSubscriber[I, O]) Request(n int64) {
	mm.sub.Request(n)
}
func (mm *MonoZipSubscriber[I, O]) Cancel() {
	mm.sub.Cancel()
}

func (mm *MonoZipSubscriber[I, O]) CurrentContext() {

}
