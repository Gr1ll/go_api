package about

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// create json types
type AboutMeRes struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

func GetAboutMe(w http.ResponseWriter, r *http.Request) {
	aboutMe := AboutMeRes{
		FirstName: "Cyril",
		LastName:  "Kurmann",
		Age:       18,
		Country:   "Switzerland",
	}
	json.NewEncoder(w).Encode(aboutMe)
	fmt.Println("Endpoint Hit: about_me")
}
