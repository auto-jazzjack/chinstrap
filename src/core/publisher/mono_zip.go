package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
	"chinstrap/core/util"
	"fmt"
)

type MonoZip[O util.All] struct {
	zipper func(...util.All) O
	monos  []core.CorePublisher[util.All]
}

func NewMonoZip2[I0 any, I1 util.All, O util.All](source1 Mono[I0], source2 Mono[I1], zipper func(I0, I1) O) Mono[O] {

	var v []util.All

	v = append(v, source1.actual)
	v = append(v, source2.actual)

	zip := &MonoZip[O]{
		monos: convert(v),
		zipper: func(a ...util.All) O {
			return zipper(a[0].(I0), a[1].(I1))
		},
	}
	return Mono[O]{
		actual: zip,
	}
}

func convert(input []util.All) []core.CorePublisher[util.All] {

	var retv []core.CorePublisher[util.All]
	for _, v := range input {
		_, ok := v.(Mono0[util.All])
		fmt.Print(ok)
		retv = append(retv, v.(Mono0[util.All]))
	}
	return retv
}

func (m *MonoZip[O]) SubscribeCore(actual core.CoreSubscriber[O]) {
	//for _, v := range m.monos {
	//	v.Subscribe(newMonoZipSubscriber(actual.(core.CoreSubscriber[any]), m.signal))
	//}
	//m.mono.actual.Subscribe()
}

func (m *MonoZip[O]) signal() {

}

func (m *MonoZip[O]) Subscribe(s reactive.Subscriber[O]) {
	Subscribe0(core.CorePublisher[O](m), s)
}

type MonoZipSubscriber[O util.All] struct {
	zipper   func(...util.All) O
	src      core.CoreSubscriber[O] //parent
	callback func()
	sub      reactive.Subscription
}

func newMonoZipSubscriber(mm core.CoreSubscriber[util.All], callback func()) reactive.Subscriber[util.All] {
	return &MonoZipSubscriber[util.All]{
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
