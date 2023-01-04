package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Maths in Golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4

	fmt.Println("Sum: ", mynumberOne+int(mynumberTwo))

	// random numbers using math
	// rand.Seed(time.Now().Unix())
	// fmt.Println("A random number", rand.Intn(10))

	// random numbers using crypto
	randVal, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("crypto rand number: ", randVal)

}
