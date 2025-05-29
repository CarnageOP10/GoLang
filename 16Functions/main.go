package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	greeter()
	result := adder(3, 5)
	fmt.Println(result)

	proRes := adder2(2,5,7,8)
	fmt.Println("pro: ", proRes)
}

func greeter()  {
	fmt.Println("hello")
}

func adder(a int, b int) int{
	return a+b
}

func adder2(values ...int) int{
	total := 0
	for i := range values{
		total+=values[i]
	}
	return total
}
