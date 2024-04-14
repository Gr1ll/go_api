package requestHandler

import (
	"log"
	"net/http"

	get "github.com/Gr1ll/go_api/requests/get"
	post "github.com/Gr1ll/go_api/requests/post"
)

func HandleRequests() {
	//get Requests
	http.HandleFunc("/get/about-me", get.GetAboutMe)

	//post Requests
	http.HandleFunc("/post/about-me", post.PostAboutMe)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
