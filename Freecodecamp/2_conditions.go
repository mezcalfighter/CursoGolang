package main

import "fmt"

func main(){
	messageLen := 10
	maxMessageLen := 20
	if messageLen > maxMessageLen{
		fmt.Println("Message sent")
	}else{
		fmt.Println("Message not sent")
	}

	// if you need a varible only for evaluation
	if length := maxMessageLen; length < 1{
		//do something
	}
}