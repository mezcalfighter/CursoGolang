package main

import (
	"fmt"
)

func main() {
	// All constants must be initialized since they can´t be changed during excecution
	const days int = 7
	var i int
	fmt.Println(i) 

	const x, y int = 17, 18
	const (
		min1 = -1
		min2 = -2
		min3 = -3
	)

	// 1.- you cannot change a constant
	const temp int = 15
	// temp = 14

	//2.-You cannot initiate a constant at runtime
	//const power = math.Pow(2,3)  // this is not a built-in function
	const l1 = len("Hello") // this is a built-in function therefore it can be done
 
	//3.- You cannot use a variable to initialize a constant
	t := 5
	fmt.Println((t))
	//const tc = t

	// Typed and untyped constants
	const a float64 = 5.1 //typed - data type specified to float64
	const b = 6.7 // untyped - data type not specified but assumed by Go

	const c float64 = a*b
	const str = "Hello " + "world in Go"

	const d = 5 > 10
	fmt.Println(d)

	// when typed constants the values must be the same
	//const x int = 5
	//const y float64 = 2.2 * x

	// when not typed
	const x2 = 5
	const y2 = 2.2 * 5
	fmt.Printf("Type: %T \n",y2)

	var i2 int = x2  // changes to int
	var j float64 = x2 //var j float64 = float64(x)
	var p byte = x2
	fmt.Println(i2,j,p)
}
