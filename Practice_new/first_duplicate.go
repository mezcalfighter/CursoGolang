package main

func FirstDuplicateValue(array []int) int {
	// Write your code here.
    hash := map[int]int{}
    for _, item := range array{
        if hash[item] == 1{
            return item
        }else{
            hash[item] = 1
        }
    }
	return -1
}