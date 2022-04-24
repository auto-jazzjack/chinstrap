package reactive

type Subscription interface {
	Request(n int64)
	Cancel()
}
