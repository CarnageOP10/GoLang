package main

import (
	"fmt"
	"net/http"
	"sync"
)

//lock the memory when 1 sub routine is working for others to write

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	websiteList := []string{
		"https://fb.com",
		"https://amazon.com",
	}

	for _, web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()//always at the end of a method
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
