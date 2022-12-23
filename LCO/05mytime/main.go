package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Golang")

	presentTime := time.Now()
	fmt.Println("Current Time: ", presentTime)

	// formatting
	fmt.Println("Current Time: ", presentTime.Format("01-02-2006 15:04:05 Monday")) // this is the standard time for formatting

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("created date: ", createdDate.Format("01-02-2006 Monday"))
}

// go env | grep GOOS
// GOOS="linux/windows/darwin" go build file.go
