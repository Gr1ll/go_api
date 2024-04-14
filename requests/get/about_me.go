package get

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// create json types
type AboutMeType struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

func GetAboutMe(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		aboutMe := AboutMeType{
			FirstName: "Cyril",
			LastName:  "Kurmann",
			Age:       18,
			Country:   "Switzerland",
		}
		json.NewEncoder(rw).Encode(aboutMe)
		fmt.Println("Endpoint Hit: about_me")
	}
}
