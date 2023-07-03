package main
import "fmt"

func CodelandUsernameValidation(str string) bool {
  if(len(str)>=4 && len(str)>=25){
    return true
  }
  return false
  // code goes here  
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(CodelandUsernameValidation("aa_"))
}