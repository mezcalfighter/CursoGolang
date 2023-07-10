package main

import (
	"fmt"
	"strings"
)

func isPalindrome(phrase string) bool {
	verifier := false
	phraseSlice := strings.Split(phrase, "")
	backSlice := []string{}
	for i := len(phraseSlice)-1; i >= 0; i-- {
		backSlice = append(backSlice, phraseSlice[i])
	}
	backPhrase := strings.Join(backSlice,"")
	if(phrase == backPhrase){
		verifier = true
	}
	return verifier
}

func main() {
	fmt.Println(isPalindrome("ama123"))
	fmt.Println(isPalindrome("ama"))
}
