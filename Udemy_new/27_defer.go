package main

import "fmt"

func foo()string{
	fmt.Println("This is foo()")
	return "I ran foo"
}

func bar(message string){
	fmt.Println(message)
	fmt.Println("This is bar()")
}

func foobar(){
	fmt.Println("This is foobar()")
}

func main(){
	defer bar(foo())
	fmt.Println("This is a string after defering bar and calling foobar()")
	defer foobar()
}