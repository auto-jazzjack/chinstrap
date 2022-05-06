package core

//import "chinstrap/core/reactive"
import (
	"chinstrap/core/reactive"
	"chinstrap/core/util"
)

type CorePublisher[T util.All] interface {
	reactive.Publisher[T]
	SubscribeCore(subscriber CoreSubscriber[T])
}
