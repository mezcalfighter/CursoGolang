package main

import (
	"fmt"
	"time"
)

func main(){
	language := "Golang"

	switch language{
	case "Python":
		fmt.Println("You're Learning Python")
		break
	case "Java":
		fmt.Println("You're Learning Java")
		break
	case "Javascript":
		fmt.Println("You're Learning Javascript")
		break
	case "Golang":
		fmt.Println("You're Learning Golang")
		break
	default:
		fmt.Println("You're Learning English")
		break
	}

	n := 5
	switch true{
	case n%2 == 0:
		fmt.Println("Even number!")
		break
	case n%2 != 0:
		fmt.Println("Odd Number!")
		break
	}

	hour := time.Now().Format(time.UnixDate)
	fmt.Println(hour)
}