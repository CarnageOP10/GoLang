package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

	// no inheritance in golang therefore no super or parent
	parth := User{"parth", "gg@gg.ai", true, 21}
	fmt.Println(parth)
	fmt.Printf("parth's details are: %+v\n", parth)
	fmt.Printf("parth's name is: %v\n", parth.Name)
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
