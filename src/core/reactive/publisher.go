package reactive

import "chinstrap/core/util"

type Publisher[T util.All] interface {
	Subscribe(s Subscriber[T])
}
