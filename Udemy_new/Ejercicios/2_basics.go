// This app is only an excercise for practice
// I'm planning to use slices - for and functions

package main

import "fmt"

var people = []string{"Carlos Perez", "Emanuel Camarena", "Luis Perez", "Luis Rodriguez", "Mario Alvarez"}
var users = []string{"username1", "username2", "username3", "username4", "username5"}
var pass = []string{"password1", "password2", "password3", "password4", "password5"}

func getName(index int) string {
	return people[index]
}

func verLogin(enteredPassword string, index int) bool {
	if index == -1 {
		return false
	}
	return enteredPassword == pass[index]
}

func findUserIndex(username string) int {
	foundUser := -1
outer:
	for index, user := range users {
		if username == user {
			foundUser = index
			break outer
		} else {
			continue
		}
	}
	return foundUser
}

func main() {
	// Login username6 - Doesn't exist
	foundUser := findUserIndex("username6")
	checkLogin := verLogin("passw0rd6",foundUser)
	if foundUser != -1 && checkLogin == true{
		returnedName := getName(foundUser)
		fmt.Printf("Welcome user: %v \n", returnedName)
	} else {
		fmt.Printf("Failed login for user %v is incorrect \n", "username6")
	}

	// Login username4 - Must return invalid login
	foundUser = findUserIndex("username4")
	checkLogin = verLogin("passw0rd6",foundUser)
	if foundUser != -1 && checkLogin == true{
		returnedName := getName(foundUser)
		fmt.Printf("Welcome user: %v \n", returnedName)
	} else {
		fmt.Printf("Failed login for user %v is incorrect \n", "username4")
	}

	// Login username4 - Must return Luis Rodriguez - Successful login
	foundUser = findUserIndex("username4")
	checkLogin = verLogin("password4",foundUser)
	if foundUser != -1 && checkLogin == true{
		returnedName := getName(foundUser)
		fmt.Printf("Welcome user: %v \n", returnedName)
	} else {
		fmt.Printf("Failed login for user %v is incorrect \n", "username4")
	}
}
