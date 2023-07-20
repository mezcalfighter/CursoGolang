package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {

	fmt.Println("f1 (goroutine) execution started")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 3)
	}
	fmt.Println("f1 execution finished")

	wg.Done()
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("f2 execution finished")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("Main execution started")
	fmt.Println("No. of CPUs: ", runtime.NumCPU())
	fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)

	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	go f1(&wg)
	fmt.Println("No. of Goroutines after running f1(): ", runtime.NumGoroutine())

	f2()
	wg.Wait()
	fmt.Println("Main execution stopped")
}
