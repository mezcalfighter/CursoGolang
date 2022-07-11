package main

import (
	"fmt"
	"models/models"
	"bufio"
	"os"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Cristian",
		LastName:  "Ortega",
	}

	fmt.Println("My firstname is:", u.FirstName)
	fmt.Print("Press 'Enter' to continue...")
  	bufio.NewReader(os.Stdin).ReadBytes('\n') 
}
