package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang pointers")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var addressOfmyNumber = &myNumber
	fmt.Println("Address Of myNumber: ", addressOfmyNumber)
	fmt.Println("Value Of myNumber: ", *addressOfmyNumber)

	*addressOfmyNumber = *addressOfmyNumber + 2
	fmt.Println("New value is: ", myNumber)

}
