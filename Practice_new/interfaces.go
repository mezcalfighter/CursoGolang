package main

import (
	"fmt"
	"math"
)

type circle struct{
	radius int
}

type rectangle struct{
	base float64
	height float64
}

type sqare struct{
	base float64
}

func (c circle) area()float64{
	return math.Pi * math.Pow(float64(c.radius),2)
}

func (r rectangle) area()float64{
	return r.base * r.height
}

func (s sqare) area()float64{
	return s.base * s.base
}

type figures interface{
	area() float64
}

func getArea(f figures){
	fmt.Println("Area: ",f.area())
}

func main(){
	getArea(rectangle{base: 12,height: 15})
	getArea(circle{radius: 12})
	getArea(sqare{base: 12})
}