package core

import "chinstrap/core/reactive"

type CoreSubscriber[T any] interface {
	reactive.Subscriber[T]
	CurrentContext() /**Context*/
	/*{
		return Context.empty();
	}*/

}
