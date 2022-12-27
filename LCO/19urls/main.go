package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=shkdhk454skhad"

func main() {
	fmt.Println("Handling URLs in GoLang")

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)
	fmt.Println("query params are: ", qparams)
	fmt.Println("coursename: ", qparams["coursename"])
	fmt.Printf("type of coursename: %T\n", qparams["coursename"])

	// create URL
	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfURL.String()
	fmt.Println("New url is: ", anotherURL)
}
