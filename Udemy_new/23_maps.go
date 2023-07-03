package main

import "fmt"

func main() {
	// Maps can be only compared to nil , if comparason sprintf to get a representation of string
	var employees map[string]string // not initialized
	people := map[string]string{"Daniel": "Friend"}
	_ = employees
	fmt.Println(people["Daniel"])

	v, ok := people["Daniel"]
	_ = v
	if ok {
		fmt.Printf("You have a %v named Carlos \n", v)
	} else {
		fmt.Println("You don't have a friend named Carlos")
	}

	// to compare maps
	map1 := fmt.Sprintf("%s", people)
	map2 := fmt.Sprintf("%s", employees)

	fmt.Printf("Comparason: %v", map1 == map2)

	// to delete
	

}
