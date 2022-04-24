package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type Mono[T any] struct {
}

type MonoImpl[T any] interface {
	SubscribeCore(actual core.CoreSubscriber[T])
	Subscribe(s reactive.Subscriber[T])
}

func (m *Mono[T]) SubscribeCore(actual core.CoreSubscriber[T]) {

}

func (m *Mono[T]) Subscribe(s reactive.Subscriber[T]) {

}
func OnAssembly[T any](source *Mono[T]) *Mono[T] {
	hook := OnEachOperatorHook
	if hook != nil {
		src := hook((*Mono[any])(source))
		return (src).(interface{}).(*Mono[T])
		//source = (Mono[T]) hook(source);
	}
	/*if (Hooks.GLOBAL_TRACE) {
		AssemblySnapshot stacktrace = new AssemblySnapshot(null, Traces.callSiteSupplierFactory.get());
		source = (Mono<T>) Hooks.addAssemblyInfo(source, stacktrace);
	}*/
	return nil
}
