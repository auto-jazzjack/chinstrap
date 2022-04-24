package reactive

type Publisher[T any] interface {
	Subscribe(s Subscriber[T])
}
