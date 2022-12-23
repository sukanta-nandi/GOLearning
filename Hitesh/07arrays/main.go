package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Coconut"
	fruitList[3] = "Banana"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length of Fruit list: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Tomato"}
	fmt.Println("veg list: ", vegList)
	fmt.Println("Length of veg list: ", len(vegList))

}
