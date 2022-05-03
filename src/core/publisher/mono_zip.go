package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type MonoZip[O any] struct {
	zipper func(...any) O
	monos  []Mono0[any]
}

func NewMonoZip2[I0 any, I1 any, O any](source1 Mono[I0], source2 Mono[I1], zipper func(I0, I1) O) Mono[O] {

	var v = []Mono0[any]{}

	v = append(v, Mono0[any](source1.actual.(Mono0[any])))
	v = append(v, Mono0[any](source2.actual.(Mono0[any])))

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

func (m *MonoZip[O]) SubscribeCore(actual core.CoreSubscriber[O]) {
	for _, v := range m.monos {
		v.Subscribe(newMonoZipSubscriber(actual, m, m.zipper))
	}
	//m.mono.actual.Subscribe()
}

func (m *MonoZip[O]) signal() {

}

func (m *MonoZip[O]) Subscribe(s reactive.Subscriber[O]) {
	Subscribe0(core.CorePublisher[O](m), s)
}

type MonoZipSubscriber[O any] struct {
	zipper    func(...any) O
	src       core.CoreSubscriber[O] //parent
	parentZip *MonoZip[O]
	sub       reactive.Subscription
}

func newMonoZipSubscriber[O any](mm core.CoreSubscriber[O], mz *MonoZip[O], zipper func(...any) O) reactive.Subscriber[O] {
	return &MonoZipSubscriber[O]{
		zipper:    zipper,
		src:       mm,
		parentZip: mz,
	}
}

func (mm *MonoZipSubscriber[O]) OnSubscribe(s reactive.Subscription) {
	mm.sub = s
	mm.src.OnSubscribe(mm)
}
func (mm *MonoZipSubscriber[O]) OnError(t error) {
	mm.src.OnError(t)
}
func (mm *MonoZipSubscriber[O]) OnNext(t O) error {
	mm.parentZip.signal()
	return nil
}

func (mm *MonoZipSubscriber[O]) OnComplete() {
	mm.src.OnComplete()
}

func (mm *MonoZipSubscriber[O]) Request(n int64) {
	mm.sub.Request(n)
}
func (mm *MonoZipSubscriber[O]) Cancel() {
	mm.sub.Cancel()
}

func (mm *MonoZipSubscriber[I, O]) CurrentContext() {

}
