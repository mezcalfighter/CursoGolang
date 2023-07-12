package main

import (
	"fmt"
	"runtime"
)

func f1(){
	fmt.Println("f1 (goroutine) execution started")
	for i:=1;  i <= 10; i++{
		fmt.Println(i)
	}
	fmt.Println("f1 execution finished")
}

func f2(){
	fmt.Println("f2 (goroutine) execution started")
	for i:=1;  i <= 3; i++{
		fmt.Println(i)
	}
	fmt.Println("f2 execution finished")
}

func main(){
	fmt.Println("Main execution started")
	fmt.Println("No. of CPUs: ",runtime.NumCPU())
	fmt.Println("No. of Goroutines: ",runtime.NumGoroutine())

	fmt.Println("OS: ",runtime.GOOS)
	fmt.Println("Arch",runtime.GOARCH)

	fmt.Println("GOMAXPROCS: ",runtime.GOMAXPROCS(0))

	go f1()
	fmt.Println("No. of Goroutines after running f1(): ",runtime.NumGoroutine())

	f2()
	fmt.Println("Main execution stopped")
}