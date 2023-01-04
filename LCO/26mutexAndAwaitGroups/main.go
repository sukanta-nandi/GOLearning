package main

import (
	"fmt"
	"sync"
)

func main() {
	println("Race Coditions in Golang")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	// can be used where we want to add read lock
	// mut.RLock() or mut.RUnlock(), read lock is quicker than write
	// rwmut := &sync.RWMutex{}

	var score = []int{0}

	// add number of goroutines to wait group
	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("scores: ", score)
}

// go run --race main.go
// exit status 66
