package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	const myurl = "http://localhost:8000/get"
	fmt.Println("Web Requests in Golang")
	//PerformGetRequest(myurl)
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(myurl string) {

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll((response.Body))
	fmt.Println(string(content))

	byteCount, _ := responseString.Write((content))
	fmt.Println("Bytecount: ", byteCount)
	fmt.Println("Response string: ", responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"course": "Let's Go with Golang",
			"price": 0,
			"platform": "LCO.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll((response.Body))
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Sukanta")
	data.Add("lastname", "Nandi")
	data.Add("email", "sukanta@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll((response.Body))
	fmt.Println(string(content))
}
