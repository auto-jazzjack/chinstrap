package main

import (
	Mono "chinstrap/core/publisher"
	"fmt"
)

func main() {
	println("hello world")

	Mono.Just("11").
		Map(func(s1 string) string {
			fmt.Printf("Hello")
			return "111"
		}).
		Subscribe0()

	v := Mono.Just(1112).
		Map(func(i int) int {
			return i + i
		}).
		Filter(func(i int) bool {
			return i > 10000
		}).
		Block()
	fmt.Println(v)

}
