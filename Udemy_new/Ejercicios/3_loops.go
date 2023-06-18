package main

import "fmt"

func main(){
	//Excercise 1: loop numbers 1 to 50 and show only the ones divisible by 7
	// Excercise 2: use continue statement
	//Exercise 3: break after the first 3
	// Excercise 4: divisible by 7 and 5 between 1 to 500
	counter := 0
	for i:=1;i<=500;i++{
		if i%7==0 && i%5==0{
			fmt.Println(i)
			counter++
			if counter == 3 {
				break
			}
		}else{
			continue
		}
	}
}