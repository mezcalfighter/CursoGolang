// Anonymous functions Go
package main

import "fmt"

// function that takes an int as an argument and returns another function that returns an int
func increment(x int) func() int {
    return func() int {
        x++
        return x
    }
}

func main(){
	nombre := "Emanuel"
	saludar := func ()  {
		fmt.Printf("Hola mi nombre es: %v \n",nombre)
	}
	saludar()
}