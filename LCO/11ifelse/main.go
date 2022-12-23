package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Not Regular user"
	} else {
		result = "Unknown user, result equals to 10"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {
	// 	fmt.Println()
	// }

}
