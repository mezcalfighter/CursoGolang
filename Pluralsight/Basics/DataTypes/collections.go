package main

import (
	"fmt"
)

func main () {
// Collections: Arrays, Slices, Maps, Structs

// Array - they must be the same type
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	slice2 = append(slice, 4)
	fmt.Println(slice2)

	//Creating slices
	newSlice := make([]int, 0)
	newSlice2 := make([]int, len(arr2))
	newSlice = append(newSlice, 1)
	newSlice2 = append(newSlice2, 1)
	fmt.Println(newSlice)
}