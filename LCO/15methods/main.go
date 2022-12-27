package main

import "fmt"

// functions inside structs are called methods

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	sukanta := User{"Sukanta", "sukanta.go@gmail.com", true, 26}
	fmt.Println("user details: ", sukanta)
	fmt.Printf("user details: %+v\n", sukanta)
	fmt.Printf("user email is: %v\n", sukanta.Email)

	fmt.Println("Print Using Methods: ")
	sukanta.GetStatus()
	sukanta.NewEmail()
}

// public struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// this passes a copy of the object
func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email Id of user: ", u.Email)
}
