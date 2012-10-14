package main

import (
	"html/template"
	"log"
	"os"
)

type Entry struct {
	Title string
	Url string
}

func main(){
	t, err := template.ParseFiles("base.html","defaults.html", "index.html")
	if err!= nil {
		log.Print(err)
	}
	values := map[string]interface{}{"entry": Entry{"mama", "http://youhe.ro"}, "other": "test"}
	t.Execute(os.Stdout, values)
}
