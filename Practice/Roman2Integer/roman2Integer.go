package main

import (
	"fmt"
)

func romanToInt(s string) int {
	var tempSlice []string
	finalResult := 0
	for index, v := range s {
		_ = index
		value := string(v)
		tempSlice = append(tempSlice, value)
	}
	for index, v := range tempSlice {
		_ = index
		if index == len(tempSlice)-1 {
			switch v {
			case "I":
				finalResult += 1
				break
			case "V":
				finalResult += 5
				break
			case "X":
				finalResult += 10
				break
			case "L":
				finalResult += 50
				break
			case "C":
				finalResult += 100
				break
			case "D":
				finalResult += 500
				break
			case "M":
				finalResult += 1000
				break
			}
		} else {
			switch v {
			case "I":
				if tempSlice[index+1] != "I" {
					finalResult -= 1
					break
				} else {
					finalResult += 1
					break
				}
			case "V":
				finalResult += 5
				break
			case "X":
				if  tempSlice[index+1] == "L" || tempSlice[index+1] == "C"{
					finalResult -= 10
					break
				} else {
					finalResult += 10
					break
				}
			case "L":
				finalResult += 50
				break
			case "C":
				if tempSlice[index+1] == "D" || tempSlice[index+1] == "M" {
					finalResult -= 100
					break
				} else {
					finalResult += 100
					break
				}
			case "D":
				finalResult += 500
				break
			case "M":
				finalResult += 1000
				break
			}
		}
	}
	return finalResult
}

func main() {
	fmt.Println(romanToInt("MCDLXXVI"))
}
