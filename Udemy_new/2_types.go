package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	// cannot change the type one declared
	//b = "Hello world"

	// cannot use different types between each other
	// sum := a + b

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
	fmt.Println(a, b)
}
