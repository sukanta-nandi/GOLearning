package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myChannel := make(chan int, 4)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// Receive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		val, isChanOpen := <-myChannel
		fmt.Println("Is Channel Open: ", isChanOpen)
		fmt.Println("Val: ", val)

		fmt.Println(<-myChannel)

		// we can use a buffer channel instead
		// fmt.Println(<-myChannel)
	}(myChannel, wg)

	// Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		close(myChannel)
		// myChannel <- 5
		// myChannel <- 6

	}(myChannel, wg)

	wg.Wait()
}
