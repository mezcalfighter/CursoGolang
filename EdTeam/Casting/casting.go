package main

import (
	"fmt"
)

func main() {
	var a uint16 = 356
	var b uint32 = 56461
	// al castear, esto no cambia el tipo de la variable en a
	c := uint32(a) + b
	fmt.Println("El valor es:", c)
	// \n es un salto de linea
	fmt.Printf("\nEl tipo es: %T", c)
}
