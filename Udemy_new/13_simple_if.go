package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("2aa")
	if err != nil {
		fmt.Printf("There was an error: %v \n", err)
	} else {
		fmt.Printf("The number parsed was %v \n", i)
	}

	if j, err := strconv.Atoi("45a"); err == nil {
		fmt.Printf("No error, the value is: %v \n", j)
	} else {
		fmt.Printf("The error is: %v \n", err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Printf("An argument is required \n")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Argument must be an integer! Error", err)
	} else {
		fmt.Printf("%v km in miles is %v \n", km, (float32(km) * 1.609))
	}
}
