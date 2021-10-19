package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Id      string `json:Id`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")
	fmt.Println("Endpoint Hit")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	muxServer()
}
