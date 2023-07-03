package main

import "fmt"

func QuestionsMarks(str string) bool {
	var characterTemp string
	var counter *int = new(int)
	for length := len(str); length > 0; length-- {
		characterTemp = str[length-1 : length]
		fmt.Println(characterTemp)
		if characterTemp == "?" {
			*counter++
		}
	}
	// code goes here
	if *counter >= 3 {
		return true
	}
	return false
}

func main() {
	newVar := "a?a6?9?"
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(QuestionsMarks(newVar))
}
