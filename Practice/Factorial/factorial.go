package main
import "fmt"

func FirstFactorial(num *int) int {
  	for num2 := *num; num2 >= 2; num2-- {
		*num = (*num) * (num2 - 1)
	}
  	return *num;
}

func main() {
  // do not modify below here, readline is our function
  // that properly reads in the input for you
  var num *int = new(int)
  *num = 8
  fmt.Println(num)
  fmt.Println(FirstFactorial(num))

}