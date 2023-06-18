package main

import "fmt"

func main() {
	// Declaracion de constantes
	//const  no pueden ser mutadas dentro de la ejecusion
	const PI float32 = 3.141592

	fmt.Println("Hola Mundo")
	fmt.Printf("Pi: %v", PI)

	// Declaracion de variables enteras
	// bases no utilizadas pero si declaras dan error, se tienen que poner _ para que vayan al garbage collector
	base := 12
	var base2 int = 12
	var base3 int
	fmt.Printf("Base: %v \n", base)
	_ = base2
	_ = base3

	//zero values
	var a int     // 0
	var b float64 // 0.0
	var c string  // ""
	var d bool    // false

	fmt.Printf("A: %v   B: %v    C:%v    D:%v \n", a, b, c, d)

	// Area de un cuadrado
	var baseA, altura int = 10, 20
	area := baseA * altura
	fmt.Printf("Area es igual a: %v \n", area)
}
