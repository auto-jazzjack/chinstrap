package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
	"chinstrap/core/util"
)

type InnerProducer[T util.All] interface {
	reactive.Subscription
	actual() core.CoreSubscriber[T]
}
