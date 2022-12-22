package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// taking input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating for food: ")

	// store the reader in a varibale

	// comma ok || err ok
	// Go does not have try catch....so treat problmes like a true flase in variable

	// input, _ (to ignore)/err
	input, _ := reader.ReadString('\n') // read till \n
	fmt.Println("User rating: ", input)
	fmt.Printf("Type of rating: %T \n", input)

}

// search for a package in go.dev
