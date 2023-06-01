package main

import (
	"fmt"
	"os"
)

func main(){
	employees:=[]string{"Carlos Perez Gonzalez","Ernesto Pelayo Martinez","Luis Barajas Hernandez","Emanuel Camarena"}
	if os.Args[1] == ""{
		fmt.Println("Tiene que ingresar el nombre de un empleado a buscar")
	}else{
		fmt.Printf("El empleado que usted desea encontrar es: %v \n",os.Args[1])
		find:
			for _, employee := range employees{
				if employee == os.Args[1]{
					fmt.Printf("El empleado %v fue encontrado \n",employee)
					break find
				}
			}
	}
}