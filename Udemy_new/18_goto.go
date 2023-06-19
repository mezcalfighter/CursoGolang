package main

import "fmt"

func main(){
	i := 1
loop:
	if i <= 5{
		fmt.Println(i)
		i++
		goto loop
	}
}