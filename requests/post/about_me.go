package post

import (
	"encoding/json"
	"log"
	"net/http"
)

// create json types
type AboutMeType struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

func PostAboutMe(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		var t AboutMeType
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Println(t.FirstName)
		http.Error(rw, "Post request received", http.StatusOK)
	}
}
