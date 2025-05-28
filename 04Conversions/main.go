package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcomeMsg := "welcome"
	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

	fmt.Printf("type of this rating is %T", input)

	// conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to rating", numRating+1)
	}


}
