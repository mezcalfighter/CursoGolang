package main

import "fmt"

func functSuma(x int, y int)int{
	return x+y
}

func returnNames(name string, last string)(string, string){
	return name, last
}

func returnCoordinates()(x,y int){
	return // x and y will be returned either way
}

func main(){
	resultadoSuma := functSuma(3,4);
	fmt.Printf("El resultado es: %d",resultadoSuma)
}