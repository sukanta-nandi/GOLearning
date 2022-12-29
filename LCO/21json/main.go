package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //aliases for key
	Price    int
	Platform string
	password string   `json:"-"` //do not show field
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling JSON in Golang")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "efg123", nil},
	}

	// package this data as JSON
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "MERN Bootcamp",
			"Price": 199,
			"Platform": "LearnCodeOnline.in",
			"tags": [
					"fullstack",
					"js"
			]
		}
	`)

	var lcoCourse course

	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where we want to add data to key value
	var myOnlineData map[string]interface{} // when you dont know the value type
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourse)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
