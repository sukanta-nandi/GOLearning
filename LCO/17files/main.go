package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "This is will be written in File"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		fmt.Println("Some error occured")
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("File length is: ", length)
	file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	// return data in byte string
	dataBytes, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n ", dataBytes)
	fmt.Println("Text data inside the file is: \n ", string(dataBytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
