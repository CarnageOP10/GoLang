package main

import (
	"fmt"
	"sync"
)

//1 thread writing, another came and want to write

func main() {

	fmt.Println("welcome")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}
	
	wg.Add(4)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("one routine")

		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("two routine")

		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("three routine")

		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("four routine")

		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
