package main

import (
	"fmt"
)

func main () {
	var firstName *string = new(string)
	*firstName = "Emanuel"
	fmt.Println(*firstName)

	secondFirstName := "Cristian"
	fmt.Println(secondFirstName)
	ptr := &secondFirstName
	
	fmt.Println(ptr, *ptr)
	secondFirstName = "Tricia"
	fmt.Println(ptr, *ptr)
}