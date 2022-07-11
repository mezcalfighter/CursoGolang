// Structs can be any type but they can't be modified when initialized
package main

import(
	"fmt"
)

func main() {
	type user struct{
		ID int
		FirstName string
		LastName string
	}
	var u user
	u.ID = 1
	u.FirstName = "Emanuel"
	u.LastName = "Camarena"
	fmt.Println(u)
	fmt.Println("My first name is", u.FirstName)

	u2 := user{ID: 1, FirstName: "Cristian", LastName: "Ortega"}
	fmt.Println(u2)
}