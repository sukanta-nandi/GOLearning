package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in Golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", servHome).Methods("GET")

	// run server
	log.Fatal(http.ListenAndServe(":3000", r))

}

func greeter() {
	fmt.Println("Hi Go User")
}

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Golang Home</h1>"))
}
