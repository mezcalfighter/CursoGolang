package main

import "fmt"

type rect struct {
	width int
	height int
}

type authenticationInfo struct{
	username string
	password string
}

func (a authenticationInfo) authenticate(enteredUsername string, enteredPassword string) bool{
	if a.username == enteredUsername && a.password == enteredPassword{
		fmt.Println("You have accessed the system")
		return true
	}else{
		fmt.Println("Username or Password is incorrect")
		return false
	}
}

func (r rect) area() int {
	return r.width * r.height
}

func main(){
	r := rect{
		width: 5,
		height: 10,
	}

	fmt.Println(r.area())

	a := authenticationInfo{
		username: "Emanuel",
		password: "12345",
	}

	accessed := a.authenticate("Carlos","34566")
	fmt.Println("El ingreso fue satisfactorio: ", accessed)
}