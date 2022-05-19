package main

import (
	"fmt"
)

func main() {
	// bool , string, numeric

	// Si no se asigna ningun valor, este comienza desde False
	var a bool = true
	fmt.Printf("Tipo: %T, Valor: %v", a, a)

	// No se pueden usar '', se tiene que usar ""
	var b string = "Valor"
	fmt.Printf("Tipo: %T, Valor: %v", b, b)

	//en el caso de los enteros el unitNUMERO DE BITS
	var c uint8 = 254
	fmt.Printf("Tipo: %T, Valor: %v", c, c)

	//unit8  es el formato, pero byte el alias para hacerlo identificable
	var d byte = 254
	fmt.Printf("Tipo: %T, Valor: %v", d, d)

	//unit32 puede recibir hasta 32 bytes
	var e uint32 = 97
	fmt.Printf("Tipo: %T, Valor: %v", e, e)

	// Aqui no guarda el a porque esta pasando el valor
	//  que tiene a en unit-code
	var f rune = 'a'
	fmt.Printf("Tipo: %T, Valor: %v", f, f)

	var g float64 = 23.64
	fmt.Printf("Tipo; %T, Valor: %v", g, g)

	// Para declarar variables que van a ser utilizadas despues pero no aun, se utiliza el operador blank osease
	_ = 1234
	var _ = 5678

	// El valor se comienza a asignar desde cero
	var h unit
	fmt.Printf("Tipo; %T, Valor: %v", h, h)
}
