package about

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
	decoder := json.NewDecoder(req.Body)
	var t AboutMeType
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.FirstName)
}
