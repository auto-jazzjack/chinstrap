package publisher

func Just[T any](v T) *Mono[T] {

	return OnAssembly((*Mono[T])(&NewMonoJust(v).Mono))

	//return &Mono[T]{}
}
