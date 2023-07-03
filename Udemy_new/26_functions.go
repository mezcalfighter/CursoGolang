package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function.")
}

func f2(a, b int) {
	fmt.Println("Sum is: ", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func f6(a, b int)(s int){
	s = a + b
	return
}

// variadic functions
func f7(nums ... int)(total int){
	total = 1
	for _, num := range nums{
		total *= num
	}
	return
}

func f8(nums ... int)(total int){
	total = 1
	for _, num := range nums{
		total += num
	}
	return
}

func main() {
	f1()
	f2(5, 7)
	f3(4, 5, 6, 4.4, 7.8, "golang")
	p := f4(5.1)
	println(p)
	intVar, intVar2 := f5(8, 9)
	fmt.Println(intVar, intVar2)
	fmt.Println(f6(1,2))
	fmt.Println(f7(1,1,1,3,5))
	nums := []int{4, 5, 6, 7, 8,9}
	fmt.Println(f8(nums...))
}
