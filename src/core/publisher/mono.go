package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type Mono[T any] struct {
	actual Mono0[T]
	//Map() Mono[any]
}

type Mono0[T any] interface {
	core.CorePublisher[T]
}

func Just[V any](v V) Mono[V] {
	return NewMonoJust(v)

}

func (m Mono[T]) Block() T {
	b := NewBlockingMonoSubscriber[T]()
	v := reactive.Subscriber[T](b)
	m.actual.Subscribe(v)
	return b.BlockingGet()
}

func (m Mono[T]) Map(consumer func(T) T) Mono[T] {
	return NewMonoMap(m, consumer)
}

func (m Mono[T]) Filter(predicate func(T) bool) Mono[T] {
	return NewMonoFilter(m, predicate)
}

func Zip2[I0 any, I1 any, O any](source1 Mono[I0], source2 Mono[I1], zipper func(I0, I1) O) Mono[O] {
	return NewMonoZip2(source1, source2, zipper)
}

func (m Mono[T]) Subscribe0() {
	m.actual.Subscribe(NewLamdaSubscriber[T](nil, nil, nil))
}

func (m Mono[T]) Subscribe(s reactive.Subscriber[T]) {
	pub := core.CorePublisher[T](m)
	Subscribe0(pub, s)
}

func (m Mono[T]) SubscribeCore(sub core.CoreSubscriber[T]) {
	panic("should not reached")
}

func Subscribe0[T any](m core.CorePublisher[T], s reactive.Subscriber[T]) {
	pub := core.CorePublisher[T](m)
	sub := s.(core.CoreSubscriber[T])
	pub.SubscribeCore(sub)
}
