package main

import (
	"fmt"
	"strings"
)

func change(a *int)*float64{
	*a = 100
	b:= 5.5

	return &b
}

func changeVar(a int){
	a = 66
}

func main(){
	x := 8
	p := &x
	fmt.Printf("Value of x before calling change() : %v \n",x)
	change(p)
	fmt.Printf("Value of x after calling change() : %v \n",x)
	fmt.Println(strings.Repeat("#",60))
	fmt.Printf("Value of x before calling changeVar() : %v \n",x)
	changeVar(x)
	fmt.Printf("Value of x after calling changeVar() : %v \n",x)
}