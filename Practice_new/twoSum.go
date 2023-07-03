package main

import "fmt"

func twoSum(slice []int, target int)(foundSlice []int){
	hash := map[int]bool{}
	// put all numbers available into slice
	for _, item := range slice{
		if hash[item] == false{
			hash[item] = true
		}else{
			continue
		}
	}
	// finds if the number exists in slice
	for _, number := range slice{
		targetNumber := target - number
		if hash[targetNumber] == true{
			foundSlice= append(foundSlice,targetNumber,number)
			break
		}
	}

	return foundSlice
}

func main(){
	slice := []int{1,3,4,5,0}
	fmt.Println(twoSum(slice,9))
}