package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello "+r.URL.Path[1:])
}

func main() {
	log.Print("test")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
