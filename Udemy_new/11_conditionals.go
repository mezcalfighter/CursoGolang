package main

import "fmt"

func main() {
	price, inStock := 100, true
	if price > 80 {
		fmt.Printf("Too expensive! \n")
	}

	if price <= 100 && inStock == true {
		fmt.Printf("Buy it! \n")
	}

	if price < 100 {
		fmt.Println("Price is cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Price is expensive")
	}
}
