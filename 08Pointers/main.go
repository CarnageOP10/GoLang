package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int
	fmt.Println("value of ptr is : ", ptr)//<nil>

	//&
	mynumber := 23
	var numpointer = &mynumber
	fmt.Println(numpointer)
	fmt.Println(*numpointer)

	*numpointer = *numpointer * 2
	fmt.Println(mynumber)



}
