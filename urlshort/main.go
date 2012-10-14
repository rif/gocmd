package main

import (
	"fmt"
	"log"
	"net/http"
)

var store = NewURLStore("database.json")

const AddForm = `
<html>
<body>
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
</body>
</html>
`

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}


func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	log.Print(url)
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}