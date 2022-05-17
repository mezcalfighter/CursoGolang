package main

import (
	"fmt"
)

const (
	pi = 3.1415
	first = iota + 6
	second
)

const (
	third = iota
)

func main () {
	fmt.Println(pi,first, second, third)
}