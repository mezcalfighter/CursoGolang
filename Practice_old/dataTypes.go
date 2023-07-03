package main

import (
	"fmt"
	"math"
)

const (
	pi = 3.1415
)

func main () {
	var radio *float64 = new(float64)
	*radio = 4
	circulo := pi * math.Pow(*radio, 2)
	fmt.Println("El área del circulo es",circulo)
}