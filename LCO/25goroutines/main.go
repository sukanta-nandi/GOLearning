package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"} // shared resource for all goroutines

var wg sync.WaitGroup
var mut sync.Mutex // pointer

func main() {
	// greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) // how many go-routines
	}

	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
