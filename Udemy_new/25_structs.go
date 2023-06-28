package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606
	fmt.Printf("Title: %v  Author: %v   Year: %v \n", title1, author1, year1)
	fmt.Printf("Title: %v  Author: %v   Year: %v \n", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	myBook := book{"The Devine Comedy", "Dante Aligheri", 1320}
	myBook2 := book{title: "The Devine Comedy", author: "Dante Aligheri", year: 1320}
	fmt.Println(myBook)
	fmt.Printf("Title: %v  Author: %v   Year: %v \n", myBook.title, myBook.author, myBook.year)
	fmt.Printf("Title: %v  Author: %v   Year: %v \n", myBook2.title, myBook2.author, myBook2.year)

	// Anonymous Struct
	// No explicitly defined struct type alias
	diana := struct {
		firstName, lastName string
		age                 int
		float64
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
}
