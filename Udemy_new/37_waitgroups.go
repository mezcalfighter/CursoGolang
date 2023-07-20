package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 4)
	fmt.Println("Finished Executing Goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Go waitGroup Tutorial")
	var wg sync.WaitGroup
	wg.Add(1) // Add amount of tasks to add to the waitgroup
	go myFunc(&wg)
	wg.Wait()
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait() // blocks until it is 0
	fmt.Println("Finished Executing my go program")
}
