package main

import (
	"fmt"
	"sort"
)

// Slices are under the hood arrays but more advanced

func main() {
	fmt.Println("Welcome to Slices in Golang")

	var fruitList = []string{"Apple", "Banana", "Coconut"} // need to initialize
	fmt.Printf("Type of fruitList: %T \n", fruitList)
	fmt.Println("Current fruitList: ", fruitList)

	// We can add values without any limit
	fruitList = append(fruitList, "Orange", "Mango")
	fmt.Println("New fruitList: ", fruitList)

	// slicing
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// use of make
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 935
	highScores[2] = 451
	highScores[3] = 874
	//highScores[4] = 777

	fmt.Println("HighScores: ", highScores)

	highScores = append(highScores, 555, 666, 777)

	fmt.Println("HighScores: ", highScores)

	// sort package
	sort.Ints(highScores)
	fmt.Println("Sorted HighScores: ", highScores)
	fmt.Println("IS Sorted HighScores: ", sort.IntsAreSorted(highScores))

	// how to remove value from slices
	var courses = []string{"ReactJS", "Python", "Ruby", "Dart"}
	fmt.Println("Courses: ", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("New Course List: ", courses)

}
