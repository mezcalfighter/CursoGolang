package main

import (
	"fmt"
	"strings"
)

func main(){
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello\"")
	//baltics =  option + }] 
	fmt.Println(`He says: "Hello"`)

	// Slicing a string - 
	s2 := "I love Golang"
	fmt.Println(s2[2:6])

	fmt.Println(strings.Contains("I love Golang Programming!","love"))
	fmt.Println(strings.ContainsAny("Hello World","eo"))
	fmt.Println(strings.Count("Hello World",""))
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.EqualFold("go","Go"))
	fmt.Println(strings.Repeat("#",30))
	fmt.Println(strings.Replace("127.0.0.1,8080",",",":",-1))
	fmt.Println(strings.ReplaceAll("127.0.0.1,8080",",",":"))
	fmt.Println(strings.Split("a,b,c,d,e",","))
}