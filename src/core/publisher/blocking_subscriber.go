package publisher

import (
	"chinstrap/core/reactive"
	"math"
	"sync"
)

type BlockingMonoSubscriberImpl[T any] struct {
	wg    sync.WaitGroup
	err   error
	value T
	sub   reactive.Subscription
}

func NewBlockingMonoSubscriber[T any]() *BlockingMonoSubscriberImpl[T] {
	_wg := sync.WaitGroup{}
	_wg.Add(1)
	return &BlockingMonoSubscriberImpl[T]{
		wg: _wg,
	}
}

func (m *BlockingMonoSubscriberImpl[T]) OnSubscribe(s reactive.Subscription) {
	m.sub = s
	m.sub.Request(math.MaxInt64)
}
func (m *BlockingMonoSubscriberImpl[T]) OnNext(t T) error {
	m.value = t
	return nil
}
func (m *BlockingMonoSubscriberImpl[T]) OnError(t error) {
	m.err = t
	m.wg.Done()
}
func (m *BlockingMonoSubscriberImpl[T]) OnComplete() {
	m.wg.Done()
}
func (m *BlockingMonoSubscriberImpl[T]) CurrentContext() {

}

func (m *BlockingMonoSubscriberImpl[T]) BlockingGet() T {
	m.wg.Wait()
	return m.value
}
