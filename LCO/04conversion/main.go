package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 10")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// type conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // string, bitsize

	// if there is some error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to user rating: ", numRating+1)
	}
}
