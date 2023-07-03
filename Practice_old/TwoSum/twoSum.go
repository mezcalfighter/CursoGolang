package main

import(
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var newNums []int
	for i := 0; i < len(nums); i++{
		for j := 1; j < len(nums); j++{
			numTemp := nums[i]
			numTemp2 := nums[j]
			fmt.Println(numTemp)
			fmt.Println(numTemp2)
			if i == j {
				continue
			}
			if(numTemp + numTemp2) == target {
				newNums = append(newNums, i)
				newNums = append(newNums, j)
				fmt.Println(newNums)
				return newNums
			}
		}
	}
	return newNums
}

func main(){
	nums := []int{2,5,5,11}
	fmt.Println(twoSum(nums,10))
}