package main

import "fmt"

func main() {
	var mySlice []int
	_ = mySlice

	nums := make([]int, 2) // <-- []int{0,0}
	fmt.Println(nums)

	cities := []string{"Tokyo","CMDX","Berlin"}
	fmt.Println(cities)

	type names []string
	friends := names{"Dan","Emanuel","Maria"}
	_ = friends

	newFriends := append(friends,"Carlos","Juan","Pepe")
	fmt.Println(newFriends)

	// Expresions 
	// a[start:stop] 
	a := []int{1,2,3,4,5}
	b := a[0:2]
	fmt.Println(b)

	testSlice := []int {1,2,3,4,5}
	newSlice := testSlice
	// This will create an exact copy and will not be dependant to testslice
	copySlice := []int{}
	copySlice = append(copySlice, testSlice[:]...)
	newcopySlice := testSlice[0:]
	testSlice[0] = 15
	fmt.Println(newSlice)
	fmt.Println(copySlice)
	fmt.Println(newcopySlice)
}