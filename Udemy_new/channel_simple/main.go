package main

import "fmt"

func f1(n int, ch chan int){
	ch <- n
}

func main(){
	c := make(chan int)

	// only for receiving 
	c1 := make(<- chan string)
	_ = c1

	// only for sending
	c2:= make(chan <- string)
	_ = c2

	go f1(10, c)
	n := <- c

	fmt.Println("Value received", n)
	fmt.Println("Exiting main...")
}