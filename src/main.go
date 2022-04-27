package main

import Mono "chinstrap/core/publisher"

func main() {
	println("hello world")

	Mono.Just("11").
		Map(func(s1, s2 string) string {
			return "111"
		})
}
