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

	t1, ok1 := source1.actual.(Mono0[any])
	if ok1 {
		panic(ok1)
	}
	v = append(v, Mono0[any](t1))

	t2, ok2 := source2.actual.(Mono0[any])
	if ok2 {
		panic(ok2)
	}
	v = append(v, t2)

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
		v.Subscribe(newMonoZipSubscriber(actual.(core.CoreSubscriber[any]), m.signal))
	}
	//m.mono.actual.Subscribe()
}

func (m *MonoZip[O]) signal() {

}

func (m *MonoZip[O]) Subscribe(s reactive.Subscriber[O]) {
	Subscribe0(core.CorePublisher[O](m), s)
}

type MonoZipSubscriber[O any] struct {
	zipper   func(...any) O
	src      core.CoreSubscriber[O] //parent
	callback func()
	sub      reactive.Subscription
}

func newMonoZipSubscriber(mm core.CoreSubscriber[any], callback func()) reactive.Subscriber[any] {
	return &MonoZipSubscriber[any]{
		src:      mm,
		callback: callback,
	}
}

func (mm *MonoZipSubscriber[any]) OnSubscribe(s reactive.Subscription) {
	mm.sub = s
	mm.src.OnSubscribe(mm)
}
func (mm *MonoZipSubscriber[any]) OnError(t error) {
	mm.src.OnError(t)
}
func (mm *MonoZipSubscriber[any]) OnNext(t any) error {
	mm.callback()
	return nil
}

func (mm *MonoZipSubscriber[any]) OnComplete() {
	mm.src.OnComplete()
}

func (mm *MonoZipSubscriber[any]) Request(n int64) {
	mm.sub.Request(n)
}
func (mm *MonoZipSubscriber[any]) Cancel() {
	mm.sub.Cancel()
}

func (mm *MonoZipSubscriber[any]) CurrentContext() {

}
