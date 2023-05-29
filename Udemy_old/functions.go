package main

import "fmt"

func suma(numero int,numero2 int){
	var total int = 0
	total = numero + numero2
	fmt.Println(total)
} 

func main(){
	suma(5,6)
}