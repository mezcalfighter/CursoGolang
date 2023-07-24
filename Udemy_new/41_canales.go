// codigo facilito

package main

import "fmt"

func main(){
	channel := make(chan string)

	go func(channel chan string){
		for{
			var name string
			fmt.Scan(&name)
			channel <- name
		}
	}(channel)
	
	msg := <- channel
	fmt.Println(msg)

}