package main

import "fmt"

func main() {
	// harcode way
	// var variableName type = value
	var v uint8 = 17
	var w int8 = 17
	var x int = 17
	var err error

	//easy way
	// variableName := value
	y := 17

	// blank identifier - to declare varibles that won't be used
	_ := 17
	var _ int = 17

	// multiple declarations
	car, cost := "BMW", 50000
	var a, b, c int

	//multiple assignment
	a, b = 10, 11
	_, _ = a, b
}
