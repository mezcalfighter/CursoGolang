package main

import (
	"fmt"
	"strings"
)

func FirstReverse(str string) string{
	var newSlice []string
	for length := len(str); length > 0; length-- {
		newSlice = append(newSlice,str[length-1 : length])
	}
	str = strings.Join(newSlice,"")
	return str
}

func main() {
	newVar := "a?a6?9?"
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstReverse(newVar))
}
