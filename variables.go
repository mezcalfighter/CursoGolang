package main

import "fmt"

func main() {
	var dog, cat = "dog" , "cat"
	fmt.Println(dog)
	fmt.Println(cat)
	// var dog, cat string = "dog" , "cat"
	
	// Sirve para crear y asignar la variable sin necesidad del uso de la palabra var
	// dog, cat := "dog" , "cat"
	cat, face := "gatito","face"
	fmt.Println(face)
	// fmt.Println(dog)
	// fmt.Println(cat)
}