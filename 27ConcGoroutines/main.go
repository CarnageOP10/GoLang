package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

//dont communicate by sharing memory;
//instead share memory by communicating

var wg sync.WaitGroup
//add done and wait
//when go routine is being is created, add it. when go routine is done -> done

func main() {
	// greeter("welcome")
	// greeter("world")

	// fmt.Println("--------------")

	// go greeter("hello")
	// greeter("world")

	// fmt.Println("---------------")

	// go greeter2("Welcome")
	// greeter2("world")

	websiteList := []string{
		"https://fb.com",
		"https://amazon.com",
	}

	for _, web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()//always at the end of a method

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}

func greeter2(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

// waitGroups
func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}
