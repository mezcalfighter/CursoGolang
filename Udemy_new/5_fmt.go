package main

import "fmt"

func main() {
	name := "Emanuel"
	fmt.Println("Hello ", name)

	a, b := 4, 6
	fmt.Println("Sum: ", a+b, "Mean value: ", (a+b)/2)

	fmt.Printf("Your age is: %d \n", 21)
	fmt.Printf("x is %d, y is %f \n", 4, 6.8)
	fmt.Printf("He says: \"Hello world\" \n")

	figure := "Circle"
	radius := 5
	pi := 3.14150

	_, _ = figure, pi
	fmt.Printf("Radius is: %d \n", radius)
	fmt.Printf("Radius is: %+d \n", radius)

	fmt.Printf("PI constant is: %f\n", pi)

	fmt.Printf("The diameter of %s with a radius of %d is %f \n", figure, radius, float64(radius)*2*pi)

	// %q for quoted string
	fmt.Printf("This is a %q", figure)

	//%v -> replaced by any type
	fmt.Printf("The diameter of %v with a radius of %v is %v \n", figure, radius, float64(radius)*2*pi)

	// %T to print type
	fmt.Printf("Figure is of type:  %T \n", figure)
	fmt.Printf("Radius is of type:  %T \n", radius)
	fmt.Printf("PI is of type:  %T \n", pi)

	// %t -> bool
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// %b -> base2
	fmt.Printf("%08b\n", 55)
}
