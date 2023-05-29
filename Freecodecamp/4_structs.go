package main

import "fmt"

type messageToSend struct{
	message string
	phoneNumber int
}

func test(m messageToSend){
	fmt.Printf("Sending message: '%s' to %v \n", m.message, m.phoneNumber)
	fmt.Println("=====================================")
}

func main(){
	test(messageToSend{
		phoneNumber: 1234567891011,
		message: "Test message",
	})

	newMessage := messageToSend{}
	newMessage.message = "New messgae to send"
	newMessage.phoneNumber = 123456789

	anonymousMessage := struct {
		Message string
		Number int64
	} {
		Message: "Hey there I am anonymous",
		Number: 123456789,
	}

	fmt.Println(anonymousMessage)

	//Embedded Struct
	type car struct{
		make string
		model string
	}
	type truck struct{
		// heredaria todo de struct car
		car
		bedSize int
	}
}