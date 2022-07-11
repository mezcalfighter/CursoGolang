// Maps are like dictionaries it might work with JSON files
package main

import(
	"fmt"
)

func main(){
	m := map[string]int{"foo":42}
	fmt.Println(m)
	fmt.Println(m["m"])

	m["foo"] = 27
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)
}