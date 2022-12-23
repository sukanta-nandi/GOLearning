package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	sukanta := User{"Sukanta", "sukanta.go@gmail.com", true, 26}
	fmt.Println("user details: ", sukanta)
	fmt.Printf("user details: %+v\n", sukanta)
	fmt.Printf("user email is: %v\n", sukanta.Email)
}

// public struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
