package home

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//create json types
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

func home(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	json.NewEncoder(w).Encode(articles)
	fmt.Println("Endpoint Hit: home")
}