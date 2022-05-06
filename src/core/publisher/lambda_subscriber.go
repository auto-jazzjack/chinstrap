package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
	"chinstrap/core/util"
	"math"
)

type LamdaImpl[T util.All] struct {
	consumer         func(T) error
	errorConsumer    func(error)
	completeConsumer func()
	//subscriptionConsumer func(reactive.Subscription)
	sub reactive.Subscription
}

func NewLamdaSubscriber[T util.All](consumer func(T) error, errorConsumer func(error), completeConsumer func()) core.CoreSubscriber[T] {
	return &LamdaImpl[T]{
		consumer:         consumer,
		completeConsumer: completeConsumer,
		errorConsumer:    errorConsumer,
	}
}
func (l *LamdaImpl[T]) Subscribe() {

}

func (l *LamdaImpl[T]) SubscribeCore(subscriber core.CoreSubscriber[T]) {

}

func (l *LamdaImpl[T]) Request(n int64) {
	l.sub.Request(n)
}
func (l *LamdaImpl[T]) Cancel() {
	l.sub.Cancel()
}
func (l *LamdaImpl[T]) OnNext(t T) error {
	if l.consumer != nil {
		err := l.consumer(t)
		if err != nil {
			l.errorConsumer(err)
		}
	}
	return nil
}

func (l *LamdaImpl[T]) OnError(t error) {
	l.errorConsumer(t)
}

func (l *LamdaImpl[T]) OnComplete() {
	if l.completeConsumer != nil {
		l.completeConsumer()
	}
	//l.OnComplete()
}

func (l *LamdaImpl[T]) OnSubscribe(sb reactive.Subscription) {
	l.sub = sb
	l.sub.Request(math.MaxInt64)

}
func (l *LamdaImpl[T]) CurrentContext() {
	//do nothing
}
