package core

import "chinstrap/core/reactive"

type CoreSubscriber[T any] interface {
	reactive.Subscriber[T]
	currentContext() /**Context*/
	/*{
		return Context.empty();
	}*/

}
