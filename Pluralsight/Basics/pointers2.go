package main

import(
	"fmt"
)

func main(){
	// & points to the allocation in memory
	number := 32
	fmt.Println(&number)

	// &points to the value allocated in memory, as seen it points to the same place in memory
	var x int = 2
	ptr := &x
	fmt.Println(&x)
	fmt.Println(ptr)
}