package main

import "fmt"

func main() {
	type speed uint
	var s1 speed = 10
	fmt.Printf("S1 type: %T \n", s1)

	var s2 uint = 20
	var resultado speed = s1 + speed(s2)
	fmt.Println(resultado)
}
