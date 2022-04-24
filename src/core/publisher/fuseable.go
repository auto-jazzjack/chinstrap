package publisher

import (
	"chinstrap/core/reactive"
	"chinstrap/core/util"
)

const NONE = 0

const SYNC = 1

const ASYNC = 2

const ANY = 3

type Fusable interface {
}

type ScalarCallable[T any] interface {
	util.Callable[T]
}

type QueueSubscription[T any] interface {
	reactive.Subscription
	RequestFusion(mode int) int
}

type SynchronousSubscription[T any] interface {
	/*@Override
	default int requestFusion(int requestedMode) {
		if ((requestedMode & Fuseable.SYNC) != 0) {
			return Fuseable.SYNC;
		}
		return NONE;
	}*/

	RequestFusion(mode int) int
}
