package main

import "fmt"

func main() {
	type newInteger8 = int8
	var number newInteger8 = 12
	fmt.Printf("Number is %v \n", number)
}
