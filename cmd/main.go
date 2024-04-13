package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", home.home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
