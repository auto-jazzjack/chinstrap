package publisher

func Just[T any](v T) *Mono[T] {
	return &Mono[T]{}
}
