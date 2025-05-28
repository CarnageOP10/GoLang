package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello to time")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	crateddate := time.Date(2020, time.December, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(crateddate)

	fmt.Println(crateddate.Format("01-02-2006 Wednesday"))
}
