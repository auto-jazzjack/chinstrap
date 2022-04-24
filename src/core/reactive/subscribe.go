package reactive

type Subscriber[T any] interface {
	OnSubscribe(s Subscription)
	OnNext(t T)
	OnError(t error)
	OnComplete()
}
