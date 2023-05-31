package main

import "fmt"

func main() {
	// array type - fixed space
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("Numbers Type: %T \n", numbers)

	//slice - can shrink or grow
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("Cities Type: %T \n", cities)

	//map - unordered (Dictionaries in python)
	balances := map[string]float64{
		"USD": 17.59,
		"EUR": 18.98,
		"YEN": 0.13,
	}
	fmt.Printf("Balances type: %T \n", balances)

	//pointer
	var x int = 2
	ptr := &x
	fmt.Printf("Pointer type: %T and value of %v \n", ptr, ptr)
}
