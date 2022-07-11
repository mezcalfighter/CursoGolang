package main

import (
	"fmt"
	"strings"
)

func main(){
	// Arrays
	newArr := [4]int{1,2,3,4}
	fmt.Println(newArr)

	newArr2 := [...]int{1,
	2,
	3,}
	fmt.Println(newArr2)

	// Slices
	var cities *[]string
	_ = cities
	//fmt.Println(*cities)

	numbers := []int{1,2,3,4}
	fmt.Println(numbers)

	nums := make([]int, 4)
	fmt.Println(nums)

	type names []string
	friends := names{"Dan","Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best friend is",myFriend)

	friends[0] = "Gabriel"
	fmt.Println("My best friend is",friends[0])

	// Este for imprime los valores ya que index guarda el valor y value la posición
	for index, value := range numbers{
		_ = value
		fmt.Println(index)
	}
	
	fmt.Println(strings.Repeat("###",15))

	// Este for imprime la posición pero nunca los valores
	for value := range numbers{
		fmt.Println(value)
	}

	// Create your own variable types with aliases
	type entero int 
	var numero2 entero = 20
	fmt.Println(numero2)
}