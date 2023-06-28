package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//var newFile *os.File
	//_ = newFile

	newFile, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
		log.Fatal(err)
	} else {
		fmt.Printf("New File was created: %v \n", newFile)
	}

	err = os.Truncate("a.txt", 0)

	file, err := os.Open("b.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(file)
	}
	file.Close()
}
