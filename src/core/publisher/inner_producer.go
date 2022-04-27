package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type InnerProducer[T any] interface {
	reactive.Subscription
	actual() core.CoreSubscriber[T]
}
