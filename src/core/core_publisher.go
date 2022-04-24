package core

//import "chinstrap/core/reactive"
import "chinstrap/core/reactive"

type CorePublisher[T any] interface {
	SubscribeCore(subscriber CoreSubscriber[T])
	Subscribe(s reactive.Subscriber[T])
}
