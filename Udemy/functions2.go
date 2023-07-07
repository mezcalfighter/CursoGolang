package main

import (
	"fmt"
	"strings"
)

func f1(){
	fmt.Println("Hello Wolrd")
}

func f2(a int, b int){
	fmt.Println("SUM: " , a+b)
}

func f3(a, b, c int, d, e, f string, g float64){
	fmt.Println(a,b,c,d,e,f,g)
}

func f4(a, b int) int{
	return a + b
}

func f5(a,b int) (int, string){
	return a + b, "No mames"
}

func f6(a []int){
	fmt.Println(a)
}

func sum(a,b int)(s int){
	s = a + b
	return s
}

// Variadic Functions
func myVariadic(a ...int){
	fmt.Printf("%T\n",a)
	fmt.Printf("%#v\n",a)
}

func myVariadic2(a ...int){
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64){
	sum := 0.
	product := 1.
	for _, v := range a{
		sum += v
		product *= v
	}
	return sum, product
}

func PersonInformation(age int, names ...string) string{
	fullName := strings.Join(names," ")
	returnString := fmt.Sprintf("Age: %d, Full Name: %s",age,fullName)
	return returnString
}

func foo(){
	fmt.Println("This is foo")
}

func bar(){
	fmt.Println("This is bar")
}

func foobar(){
	fmt.Println("This is foobar")
}

func increment(x int) func() int{
	return func() int{
		x++
		return x
	}
}

func main(){
	f1()
	f2(2,7)
	f3(1,2,3,"No","Rompas","mas",69.99)
	fmt.Println(f4(1,4))
	sum, message := f5(1,1)
	fmt.Println("Function 5 returns: ", sum, message)
	// Multiple values - Variadic
	myVariadic(1,2,3,4)
	myVariadic()

	//example of a variadic built in
	nums := []int{1,2,3,4}
	nums = append(nums, 5,6,7)
	// It is possible to pass a slice as well
	myVariadic(nums...)
	// myVariadic2 modifies the nums slice and no copy is created
	myVariadic2(nums...)
	fmt.Println("Nums",nums)

	s, p := SumAndProduct(2.0,5.,10.,5.6,87.3)
	fmt.Println("Sum:",s,"Product:",p)

	info := PersonInformation(30,"Wolfgang","Amadeus","Mozart")
	fmt.Println(info)

	// Defer Statement - Postpose the excecution until the surrounding function returns
	defer foo()
	bar()

	fmt.Println("Just a string after defering foo() and calling bar()")
	defer foobar()

	// Anonymous Function
	func(msg string){
		fmt.Println(msg)
	}("I'm an anounymous function!")

	// Real world example for an anonymous function
	a := increment(10)
	a()
	fmt.Println(a())
}


