package main

import "fmt"

func main() {
	fmt.Println("Defer in Golang")

	defer fmt.Println("User")
	defer fmt.Println("New!")
	defer fmt.Println("World!")
	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
