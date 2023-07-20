package main

import(
	"fmt"
)

func main(){
	type book struct{
		title string
		author string
		year int
	}

	type book1 struct{
		title, author string
		year int
	}

	myBook := book{"The Devine Comedy","Dante Aligheri",1320}
	myBook1 := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	aBook := book{title: "Random book"}
	fmt.Println(myBook)
	fmt.Println(myBook1)
	fmt.Println(aBook)
	lastBook := myBook.title
	fmt.Println(lastBook)
	myBook.title = "New Book Test"
	fmt.Println(myBook)

	//Anonymous Structs

	diana := struct{
		firstName, lastName string
		age int
	}{
		firstName: "Diana",
		lastName: "Mitchel",
		age: 30,
	}
	_ = diana

	type Book struct{
		string
		float64
		bool
	}

	b1 := Book{"New Book",20.24,false}
	_ = b1

	type Employee1 struct{
		name string
		salary int
		bool
	}

	person1 := Employee1{name: "Juan Perez", salary: 1000, bool: false}
	_ = person1

	// Embedded Structs

	type Contact struct{
		email, address string
		phone int
	}

	type Employee struct{
		name string
		salary int
		contactInfo Contact
	}

	person := Employee{
		name: "Juan Perez",
		salary: 1000,
		contactInfo: Contact{
			email: "juan@email.com",
			address: "123 Pearson St",
			phone: 3333333,
		},
	}

	fmt.Println(person)
	fmt.Println("Juan email's is: " + person.contactInfo.email)

}