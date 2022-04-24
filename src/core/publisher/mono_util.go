package publisher

func Just[T any](v T) *Wrapper[T] {

	//return OnAssembly((*Mono[T])(&NewMonoJust(v).Mono))
	return NewWrapper[T](
		NewMonoJust(v),
	)

	//return &Mono[T]{}
}
