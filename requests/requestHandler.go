package requestHandler

import (
	"log"
	"net/http"

	homeGet "github.com/Gr1ll/go_api/requests/get"
)

func HandleRequests() {
	http.HandleFunc("/home", homeGet.Home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}