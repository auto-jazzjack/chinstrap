package publisher

import (
	"chinstrap/core/reactive"
)

var OnEachOperatorHook func(reactive.Publisher[interface{}]) reactive.Publisher[interface{}]
