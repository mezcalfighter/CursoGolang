package main

import (
	"fmt"
)

func main(){
	var employees map[string]string
	fmt.Println(employees)

	var accounts map[string]float64
	fmt.Printf("%v\n",accounts["0x323"])

	//employees["Dan"] = "Programmer"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8
	fmt.Println(map1)

	balances := map[string]float64{
		"USD":21.25,
		"EUR":23.00,
		"MXN":0.89, 
	}

	tempVar := balances["USD"]

	fmt.Println(balances)
	fmt.Println(tempVar)
}