package core

import "chinstrap/core/reactive"

type CoreSubscriber[T any] interface {
	currentContext() /**Context*/
	/*{
		return Context.empty();
	}*/

	OnSubscribe(s reactive.Subscription)
	OnNext(t T)
	OnError(t error)
	OnComplete()
}
