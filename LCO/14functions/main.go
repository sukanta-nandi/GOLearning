package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo() // declaring functions inside another func not allowed

	result := adder(3, 6)
	fmt.Println("Adder Result: ", result)

	proRes, msg := proAdder(2, 5, 4, 8, 9)
	fmt.Println("Pro Adder Result: ", proRes)
	fmt.Println("Pro Msg: ", msg)

}

func greeter() {
	fmt.Println("Hello from Golang")
}

func greeterTwo() {
	fmt.Println("Calling another method")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

// it is possible to return multiple values
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Test String"
}
