package main

import "fmt"

func main() {
	var mv int        // max value
	var max_value int // NOT OKAY

	var packetsReceived int // NOT OKAY, name too long
	var b int               // ok

	const MAX_VALUE = 100 // NOT OKAY
	const N = 1           // Better

	maxValue := 10000   // recommended
	max_values := 10000 // Not recommended

	writeToDB := true  // ok, idiomatic
	writeToDb := false // not ok
}
