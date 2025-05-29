package main

import "fmt"

// defer puts that line at last
// works on 'LIFO'
func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
	mydiffer()
	mydiffer2()
}


func mydiffer(){
	for i := 0; i < 5; i++{
		fmt.Println("hello: ", i)
	}
}

func mydiffer2(){
	for i := 0; i < 5; i++{
			defer fmt.Println("hello: ", i)
	}
}