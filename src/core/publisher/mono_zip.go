package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoZip[O any] struct {
	zipper func(...any) O
	monos  []Mono[any]
}

func NewMonoZip2[I0 any, I1 any, O any](source1 Mono[I0], source2 Mono[I1], zipper func(I0, I1) O) Mono[O] {

	var v = []Mono[any]{}

	v = append(v, Mono[any](source1))
	v = append(v, Mono[any](source2))

	zip := &MonoZip[O]{
		monos: v,
		zipper: func(a ...any) O {
			return zipper(a[0].(I0), a[1].(I1))
		},
	}
	return Mono[O]{
		actual: zip,
	}
}

func (m *MonoZip[I, O]) SubscribeCore(actual core.CoreSubscriber[O]) {
	for _, v := range m.monos {
		v.actual.Subscribe(newMonoZipSubscriber(actual, m, m.zipper))
	}
	//m.mono.actual.Subscribe()
}

func (m *MonoZip[I, O]) signal() {

}

func (m *MonoZip[I, O]) Subscribe(s reactive.Subscriber[O]) {
	Subscribe0(core.CorePublisher[O](m), s)
}

type MonoZipSubscriber[I any, O any] struct {
	zipper    func(...any) O
	src       core.CoreSubscriber[O] //parent
	parentZip *MonoZip[I, O]
	sub       reactive.Subscription
}

func newMonoZipSubscriber[I any, O any](mm core.CoreSubscriber[O], mz *MonoZip[I, O], zipper func(...any) O) reactive.Subscriber[I] {
	return &MonoZipSubscriber[I, O]{
		zipper:    zipper,
		src:       mm,
		parentZip: mz,
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
	return mm.src.OnNext(t)
}

func (mm *MonoZipSubscriber[I, O]) OnComplete() {
	mm.src.OnComplete()
	mm.parentZip.signal()
}

func (mm *MonoZipSubscriber[I, O]) Request(n int64) {
	mm.sub.Request(n)
}
func (mm *MonoZipSubscriber[I, O]) Cancel() {
	mm.sub.Cancel()
}

func (mm *MonoZipSubscriber[I, O]) CurrentContext() {

}
