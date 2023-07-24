package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, ch chan string){
	resp, err := http.Get(url)
	if err != nil{
		//fmt.Println(err)
		s := fmt.Sprintf("%s is Down \n",url)
		s += fmt.Sprintf("Error: %v \n",err)
		ch <- s //sending into the channel
	}else{
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200{
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split("url", "//")[0]
			file += ".txt"
			s += fmt.Sprintf("Writing response body to %s \n",file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil{
				// log.Fatal(err)
				s += "Error writing file \n"
				ch <- s
			}
			s += fmt.Sprintf("%s is UP \n",url)
			ch <- s
		}
	}
}

func main(){
	urls := []string{"http://golang1.org", "https://www.google.com/a.html", "https://youtube.com", "https://twitter.com"}

	// 1. 

	c := make(chan string)

	for _, url := range urls{
		go checkAndSaveBody(url,c)
		fmt.Println(strings.Repeat("#",20))
	}

	fmt.Println("No. of Goroutines:",runtime.NumGoroutine())

	for i:= 0; i < len(urls); i++{
		fmt.Println(<- c)
	}
}