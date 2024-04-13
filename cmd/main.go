package main

import (
	"log"
	"net/http"
	home "./get"
)

func handleRequests() {
	http.HandleFunc("/", home.home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
