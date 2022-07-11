package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Hello world")
	port := 3000
	_, err := startWebServer(port,3)
	fmt.Println(err)
}

func startWebServer(port, numberOfRetries int) (int, error) /* Here goes the return type*/{
	fmt.Println("Starting web server...")
	// doing important things
	fmt.Println("Server started on port",port)
	return port, errors.New("Something went wrong")
}
