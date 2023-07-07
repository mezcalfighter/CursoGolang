package main

import (
	"fmt"
	"math"
)

// interface
type shape interface{
	getArea() float64
	getPerimeter() float64
}

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (c circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) getArea() float64 {
	return r.height * r.width
}

func (r rectangle) getPerimeter() float64 {
	return (2 * r.height) + (2 * r.width)
}

// NOT RECOMMENDED - USE INTERFACES

func printCircle(c circle) {
	fmt.Println("Shape: ", c)
	fmt.Println("Area: ", c.getArea())
	fmt.Println("Perimeter", c.getPerimeter())
}

func printRectangle(r rectangle) {
	fmt.Println("Shape: ", r)
	fmt.Println("Area: ", r.getArea())
	fmt.Println("Perimeter", r.getPerimeter())
}

// INTERFACES - Behavior that must be implemented for similar type of objects.
func print(s shape){
	fmt.Println("Shape: ", s)
	fmt.Println("Area: ", s.getArea())
	fmt.Println("Perimeter", s.getPerimeter())
}


func main() {
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3., height: 2.1}

	printCircle(c1)
	printRectangle(r1)
}
