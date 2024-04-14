package main

import (
	"log"
	"net/http"

	homeGet "github.com/Gr1ll/go_api/requets/get"
)

func handleRequests() {
	http.HandleFunc("/", homeGet.Home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
