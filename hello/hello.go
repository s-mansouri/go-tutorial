package main

import (
	"fmt"
	"greetings"
	"rsc.io/quote"
)

func main() {

	name := "Bob"
	fmt.Println("hello world", name)
	fmt.Println(quote.Go())

	fmt.Println(greetings.Hello(name))
	const number = "90"

	var s string
	s = number

	fmt.Println("number =", number, "s =", s)

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			fmt.Println("even numbers: ", n)
		} else {
			fmt.Println("odd numbers: ", n)
		}
	}
}
