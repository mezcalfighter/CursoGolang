package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	//harcode type
	var numero1 int8 = 12
	var numero2 rune = 8
	var resultado rune = rune(numero1) + numero2

	// easy way
	numeroUno := 12
	numeroDos := 8
	resultadoNumeros := numeroUno + numeroDos

	fmt.Println("El resultado es: ",resultado)
	fmt.Println("El resultado de la suma es: ", resultadoNumeros)

	// multiple
	age, name := 26, "Emanuel"
	fmt.Println("I am ", name, "and I am ", age, " years old")

	//constants
	const secInMin int8 = 60
	const minInHour int8 = 60
	const secInHour int16 = int16(secInMin) * int16(minInHour)
	fmt.Println(secInHour)

	msg := fmt.Sprintf("Mi nombre es %v ","Emanuel")
	fmt.Printf(msg)
}
