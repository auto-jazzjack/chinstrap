package util

type Callable[T any] interface {
	Call() (T, error)
}
