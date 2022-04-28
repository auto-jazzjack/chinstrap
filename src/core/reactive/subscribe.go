package reactive

type Subscriber[T any] interface {
	OnSubscribe(s Subscription)
	OnNext(t T) error
	OnError(t error)
	OnComplete()
}
