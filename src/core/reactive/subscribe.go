package reactive

import "chinstrap/core/util"

type Subscriber[T util.All] interface {
	OnSubscribe(s Subscription)
	OnNext(t T) error
	OnError(t error)
	OnComplete()
}
