package publisher

import (
	"chinstrap/core"
	"chinstrap/core/reactive"
)

type Wrapper[T any] struct {
	internal reactive.Publisher[T]
}

type WrapperImpl[T any] interface {
	Mono[T]
}

func NewWrapper[T any](m Mono[T]) *Wrapper[T] {
	return &Wrapper[T]{
		internal: m,
	}
}

func (w *Wrapper[T]) GetWrapper() *Wrapper[T] {
	return w
}

func (w *Wrapper[T]) SubscribeCore(actual core.CoreSubscriber[T]) {

}

func (w *Wrapper[T]) Subscribe(s reactive.Subscriber[T]) {
	panic("not implemented") // TODO: Implement
}

func (w *Wrapper[T]) Map() Mono[any] {
	panic("not implemented") // TODO: Implement
}
