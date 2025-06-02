package main

import (
	"fmt"
	"sync"
)

//channels->communication between go routines

func main() {
	fmt.Println("welcome")
	
	// myCh := make(chan int)
	myCh := make(chan int, 2)//add buffer channel
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch <- chan int, wg *sync.WaitGroup){//outside the channel
		
		val, isChannelOpen := <-myCh//to check if channel is open/not :0 example:
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		
		fmt.Println(<-myCh)
		
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup){//inside the channel
		
		myCh <- 5
		myCh <- 6
		close(myCh)

		wg.Done()
	}(myCh, wg)

	wg.Wait()
	
}
