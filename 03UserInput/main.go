package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "welcome"
	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the reading for our pizza:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

	fmt.Printf("type of this rating is %T", input)


}
