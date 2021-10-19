package main

import (
	"log"
	"net/http"
)

func simpleServer() {
	Articles = []Article{
		Article{Title: "title1", Desc: "desc1", Content: "content1"},
		Article{Title: "title2", Desc: "desc2", Content: "content2"},
	}
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
