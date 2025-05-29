package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to if else")

	loginCount := 23

	var result string

	if loginCount < 23 {
		result = "less"
	} else if loginCount > 23{
		result = "greater"
	} else{
		result = "equal"
	}

	fmt.Println(result)


	if 9%2 == 0{
		fmt.Println("number is even")
	} else{
		fmt.Println("number is odd")
	}

	//assign and check on the go
	if num := 3; num < 10{
		fmt.Println("less than 10")
	} else{
		fmt.Println("num is not less than 10")
	}
}
