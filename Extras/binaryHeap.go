/* Binary Heap: Is a data structure that takes 
                the form of a binary tree*/

// Each node has a max of 2 child nodes
// The Binary heap must have the following properties... 

package main

import(
    "fmt"
)

type newArr struct{
    number []int
}

func childFinderLeft(n int) int{
    return (n*2)+1
}

func childFinderRight(n int) int{
    return (n*2)+2
}

func parentFinder(n int) int{
    return (n-1)/2
}

// Missing DATA AND WAYS TO ORDER SLICE

func main(){

}