package publisher

import "chinstrap/core"

type ScalarSubscriptionImpl[T any] interface {
	InnerProducer[T]
	SynchronousSubscription[T]
}
type ScalarSubscription[T any] struct {
	actual core.CoreSubscriber[T]

	value T

	stepName string

	once int
}

func NewScalarSubscription[T any](actual core.CoreSubscriber[T], value T) *ScalarSubscription[T] {
	return &ScalarSubscription[T]{
		actual:   actual,
		value:    value,
		stepName: "",
	}
}

func (s *ScalarSubscription[T]) Request(n int64) {

}

func (s *ScalarSubscription[T]) Cancel() {

}

func (s *ScalarSubscription[T]) Clear() {

}

func (s *ScalarSubscription[T]) IsEmpty() {

}

func (s *ScalarSubscription[T]) Acutual() {

}
