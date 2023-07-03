package main

import "fmt"

func main(){
	/*
		&value => pointer
		*pointer => value
		*type => declaring a pointer
	*/


	// & -> Address of operator
	name := "Emanuel"
	fmt.Println(&name) // This will show the hexadecimal memory allocation of name

	// Type: *int -> pointer to int
	var x int = 2
	ptr := &x
	fmt.Printf("This variable is value of: %T  with value of %v \n",ptr,ptr)

	// Create a *float64 -> Pointer to float64 it is a zero value is nil
	var b *float64
	_ = b
	// or
	p := new(int)
	x2 := 100
	p = &x2
	fmt.Printf("P is a point to x2 in address: %v \n",p)

	// * -> derefencing operator
	*p = 90
	// x2 changed value and p changed the pointer to a different value
	// we used *p to get the value that the pointer is referencing to
	fmt.Println(x2, *p)
	fmt.Printf("X2 == *p : %v \n", x2 == *p)


}