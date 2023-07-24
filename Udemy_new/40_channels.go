package main

import "fmt"

func main(){
	// nil (not initialized) channel
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch) // Pointer to channel

	// create an initialized channel
	c := make(chan int)

	// <- channel operator

	// send operation
	c <- 10

	// receive operation
	num := <- c
	fmt.Println(<- c)
	_ = num

	//close channel
	close(c)
}