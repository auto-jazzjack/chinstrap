package core

//import "chinstrap/core/reactive"
import "chinstrap/core/reactive"

type CorePublisher[T any] interface {
	reactive.Publisher[T]
	SubscribeCore(subscriber CoreSubscriber[T])
}
