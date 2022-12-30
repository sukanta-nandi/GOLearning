package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sukanta-nandi/mongoapi/routers"
)

func main() {
	fmt.Println("Simple MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started ...")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening to port 3000")
}
