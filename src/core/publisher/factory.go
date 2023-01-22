package publisher

import (
	"chinstrap/core"
	"chinstrap/core/util"
)

func Convert[T util.All](i any) core.CorePublisher[T] {

	return nil
	/* v, ok := i.(core.CorePublisher[int])
	if ok {
		return v
	} */
}

/* type InnerPublisher[T util.All] struct {
	m Mono0[T]
}

func (publisher InnerPublisher[T]) SubscribeCore(subscriber core.CoreSubscriber[util.All]) {
	//publisher.m.SubscribeCore(subscriber)
}

func (publisher InnerPublisher[T]) Subscribe(s reactive.Subscriber[util.All]) {
	//v := reactive.Subscriber[T](s)
	//publisher.m.Subscribe(v)
	publisher.m.Subscribe(s.(reactive.Subscriber[T]))
}
func NewCorePublisher[T util.All](m Mono[T]) core.CorePublisher[util.All] {

	return &InnerPublisher[T]{
		m: m.actual,
	}
}
*/
/*
type InnerSubscriber[T util.All] struct {
}

func NewCoreSubscriber[T util.All](input util.All) core.CoreSubscriber[util.All] {
	return &InnerSubscriber[util.All]{}
}
func (sub *InnerSubscriber[T]) CurrentContext() {

}
func (sub *InnerSubscriber[T]) OnSubscribe(s reactive.Subscription) {

}
func (sub *InnerSubscriber[T]) OnNext(t T) error {
	return nil
}
func (sub *InnerSubscriber[T]) OnError(t error) {

}
func (sub *InnerSubscriber[T]) OnComplete() {

} */

/* func (publisher InnerPublisher[T]) SubscribeCore(subscriber core.CoreSubscriber[util.All]) {
	publisher.m.SubscribeCore(subscriber.(core.CoreSubscriber[T]))
}

func (publisher InnerPublisher[T]) Subscribe(s reactive.Subscriber[util.All]) {
	publisher.m.Subscribe(s.(reactive.Subscriber[T]))
	//publisher.m.Subscribe(s)
}
func ConvertToObject[T util.All](m Mono[T]) core.CorePublisher[util.All] {
	return &InnerPublisher[T]{
		m: m.actual.(core.CorePublisher[T]),
	}
} */
