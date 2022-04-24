package publisher

func Just[T any](v T) *Mono[T] {

	return OnAssembly()

	//return &Mono[T]{}
}
