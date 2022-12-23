package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("Langunage Maps: ", languages)

	fmt.Println("JS shorts for: ", languages["JS"])

	// add value jus use key value mapping

	// delete a value
	delete(languages, "RB")

	fmt.Println("New Langunage Map: ", languages)

	// looping through maps
	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
