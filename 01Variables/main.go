package main

import "fmt"

func main() {
	fmt.Println("variables")
	var username string = "parth"
	fmt.Println(username)
	fmt.Println("parth fp")
	fmt.Printf("variable is of type: %T \n", username)
// bool, uint#, 

	//implicit way
	var website = "parth singh"
	fmt.Println(website)
	// website = 3 throws error

	// no var walrus operator but cant use this operator outside the function
	name := "hello"
	fmt.Println(name)
	
}
