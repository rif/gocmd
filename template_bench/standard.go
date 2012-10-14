package main

import (
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("standard/base.html", "standard/index.html")
	t.Execute(w, "hello world")
}

func main() {
	http.HandleFunc("/", mainPage)
	log.Print("serving at 8000...")
	http.ListenAndServe(":8000", nil)
}
