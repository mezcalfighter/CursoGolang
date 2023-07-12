package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){
	go mi_nombre_lento("Emanuel",2000)
	go mi_nombre_lento("Carlos",1000)
	fmt.Println("Que aburrido")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lento(name string, timeWait int){
	letras := strings.Split(name,"")
	for _, letra := range letras{
		time.Sleep(time.Duration(timeWait) * time.Millisecond)
		fmt.Println(letra)
	}
}