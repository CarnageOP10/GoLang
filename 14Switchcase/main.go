package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1

	fmt.Println("value is: ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1 and u can open")
	case 2:
		fmt.Println("move 2 spots")
	case 3:
		fmt.Println("move 3 spots")
	case 4:
		fmt.Println("move 4 spots")
	case 5:
		fmt.Println("move 5 spots")
	case 6:
		fmt.Println("dice value is 6 and u can open")	
	}

	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1 and u can open")
	case 2:
		fmt.Println("move 2 spots")
	case 3:
		fmt.Println("move 3 spots")
	case 4:
		fmt.Println("move 4 spots")
		fallthrough
	case 5:
		fmt.Println("move 5 spots")
		fallthrough
	case 6:
		fmt.Println("dice value is 6 and u can open")	
	}
}
