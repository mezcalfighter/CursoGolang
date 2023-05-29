// Las interfaces son metodos que pueden aplicar a distintos structs

package main

import "fmt"

type figuras2D interface{
	area() float64
}

type cuadrado struct{
	base float64
}

type rectangulo struct{
	base float64
	altura float64
}

func (c cuadrado) area() float64{
	return c.base * c.base
}

func (r rectangulo) area() float64{
	return r.base * r.altura
}

func calcular(f figuras2D){
	fmt.Println("Area", f.area())
}

func main(){
	miCuadrado := cuadrado{base: 2}
	miRectangulo := rectangulo{base: 4,altura: 5}
	calcular(miCuadrado)
	calcular(miRectangulo)
}