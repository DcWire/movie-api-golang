package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DcWire/mymodules/router"
)

func main() {
	fmt.Println("Mongo db API")
	router := router.Router()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my API")
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
