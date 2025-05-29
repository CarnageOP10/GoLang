package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")
	var vegList = []string{}
	fmt.Printf("TYPE IS %T \n", vegList)

	vegList = append(vegList, "aple", "mngo")
	fmt.Println(vegList)

	vegList = append(vegList[1:])
	fmt.Println(vegList)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 235
	highscores[2] = 233
	highscores[3] = 432
	//highscores[4] = 98 //will crash due to highscores := make([]int, 4)

	highscores = append(highscores, 555, 666, 777) //this wont give error
	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// removing based on index

	var courses = []string{"react", "ruby", "python"}
	fmt.Println(courses)

	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
