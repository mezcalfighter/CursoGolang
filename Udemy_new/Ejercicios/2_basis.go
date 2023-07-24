package main

import "fmt"

func login(enteredUsername string, enteredPassword string)bool{
	return true
}

func main(){
	// Simulates 2 failed loggin attemps. It will use a function to verify the login
	// using a foor loop to replicate it in 3 attempts 

	//dn_sim simulates a db but in reality we wouldn't use an array
	db_sim := [4]string{"username4","password4","username2","password2"}

	// attemps simulates de username and password entered but it would most likely be with a struct
	attempts := [6]string{"username1","password1","username2","password2","username3","password3"}

	
}