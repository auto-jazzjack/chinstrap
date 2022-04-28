package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

/* type ScalarSubscription[T any] interface {
	core.CoreSubscriber[T]
	reactive.Subscription
	//SynchronousSubscription[T]
} */
type ScalarSubscriptionImpl[T any] struct {
	actual core.CoreSubscriber[T]
	value  T
	sub    reactive.Subscription
}

func NewScalarSubscription[T any](actual core.CoreSubscriber[T], value T) reactive.Subscription {
	return &ScalarSubscriptionImpl[T]{
		actual: actual,
		value:  value,
	}
}

func (s *ScalarSubscriptionImpl[T]) Request(n int64) {
	s.actual.OnNext(s.value)
	s.OnComplete()
}

func (s *ScalarSubscriptionImpl[T]) Cancel() {
	s.sub.Cancel()
}

func (s *ScalarSubscriptionImpl[T]) OnNext(t T) error {
	return s.actual.OnNext(s.value)
}

func (s *ScalarSubscriptionImpl[T]) OnError(t error) {
	s.actual.OnError(t)
}

func (s *ScalarSubscriptionImpl[T]) OnComplete() {
	s.actual.OnComplete()
}

func (s *ScalarSubscriptionImpl[T]) OnSubscribe(sb reactive.Subscription) {
	s.actual.OnSubscribe(s)
	s.sub = sb
}

func (s *ScalarSubscriptionImpl[T]) CurrentContext() {
	//do nothing
}
