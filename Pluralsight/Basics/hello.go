package main

import (
	// Can't have libraries not used
	"fmt"
)

func main()  {
	// Can't have a variable that's not being used
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Emanuel"
	fmt.Println("Hello", firstName)

	h := true
	fmt.Println(h)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c) , imag(c)
	fmt.Println(r, im)

}