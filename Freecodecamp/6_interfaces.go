package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message){
	fmt.Println(msg.getMessage())
}

type message interface{
	getMessage() string
}

type birthdayMessage struct{
	birthdayTime time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string{
	return fmt.Sprintf("Hi %s, it is your birthday on %s",bm.recipientName,bm.birthdayTime)
}

func main(){

}