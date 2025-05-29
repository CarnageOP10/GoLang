package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "iceapple"
	// fruitList[2] = "custardapple"  this gives a blank space
	fruitList[3] = "pineapple"

	fmt.Println("fruitlist is: ", fruitList)
	fmt.Println("length is: ", len(fruitList))//gives the predefined size

	var vegList = [3]string{"bhindi","beans","potato"}
	fmt.Println(vegList, "\n", len(vegList))
}
