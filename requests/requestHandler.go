package requestHandler

import (
	"log"
	"net/http"

	about_me "github.com/Gr1ll/go_api/requests/get"
)

func HandleRequests() {
	//get Requests
	http.HandleFunc("/about-me", about_me.GetAboutMe)

	log.Fatal(http.ListenAndServe(":10000", nil))
}
