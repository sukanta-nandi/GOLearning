package main

import (
	"fmt"
)

// constants , Capital first => Public
const LoginToken string = "shdsahdhsakdhkhsd"

func main() {
	fmt.Println("Variables")

	// variables
	var username string = "sukanta"
	fmt.Println(username)

	// use of placeholder (%T)
	fmt.Printf("Varibale is of type: %T \n", username)

	// bool
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varibale is of type: %T \n", isLoggedIn)

	// Int
	var smallVal uint8 = 255 // 256 gives error : 2^8-1
	fmt.Println(smallVal)
	fmt.Printf("Varibale is of type: %T \n", smallVal)

	// other int types 16, 32, 64
	// can use int in general int == unint 32 or 64
	// same for float and float 32

	// Float
	var smallFloat float32 = 255.4545454544545454 // 256 gives error : 2^8-1
	fmt.Println(smallFloat)
	fmt.Printf("Varibale is of type: %T \n", smallVal)

	var floatVal float64 = 255.54545454544545454
	fmt.Println(floatVal)
	fmt.Printf("Varibale is of type: %T \n", floatVal)

	// default values and aliases
	var anotherVariable int // declaration
	fmt.Println((anotherVariable))
	fmt.Printf("Varibale is of type: %T \n", anotherVariable)

	// implicit type : Lexer decides var type. You can not change var type later
	var website = "learning go"
	fmt.Println(website)
	fmt.Printf("Varibale is of type: %T \n", website)

	// no var style, only allowed inside a method
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varibale is of type: %T \n", LoginToken)

}
