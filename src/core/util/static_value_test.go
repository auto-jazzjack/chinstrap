package util

import (
	"fmt"
	"testing"
)

type MyType interface {
	any | int64 | float64 | int | string | int32
}

type Temp[T MyType] interface {
	hello(v T)
}

type Templ1[T MyType] struct {
}

type TemlVoid struct {
	real Temp[MyType]
}

func (te TemlVoid) hello(v MyType) {
	te.real.hello(v)
}

func (te Templ1[T]) hello(v T) {
	fmt.Print("tempV2")
}

func TestXxx(t *testing.T) {
	b := Temp[MyType](Templ1[string]{})
	b.hello("22")
}

func create[T All](tmp Templ1[T]) TemlVoid {
	return TemlVoid{
		real: tmp,
	}
}
