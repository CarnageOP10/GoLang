package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in go lang")

	days := []string{"mon", "tue", "wed", "thurs"}

	// like c++, java no ++i here
	for day := 0; day < len(days); day++ {
		fmt.Println(days[day])
	}

	//easier
	for i := range days {
		fmt.Println(days[i])
	}

	//like java
	for index, day := range days {
		fmt.Println(day, index)
	}

	// break
	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		// if rougeValue == 5{
		// 	break
		// }
		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println("value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("jumping at some place")
}
