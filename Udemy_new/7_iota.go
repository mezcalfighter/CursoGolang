// It is a constant declaration and represents successive untyped integer constant.

package main

import "fmt"

func main(){
	const (
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)

	const (
		c00 = iota
		c11
		c22
	)

	const(
		North = iota // default 0
		East // 1
		South // 2
		West // 3
	)

	fmt.Println(West)

	const (
		a = iota * 2 // 0  * 2 = 0
		b            // 1 * 2 = 2
		c            // 2* 2 = 4
	)

	// to skip values
	const (
		x = -(iota + 2) // -2
		_				// -3
		y				// -4
		z				// -5
	)
}