package main

import (
	"log"
	"fmt"
	"strings"
	"time"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {
	log.Print(req)
	fmt.Fprint(w, strings.Repeat(" ",1200))
	fmt.Fprint(w, "<html><body><h1>Hi")
	time.Sleep(5*10e9)
	fmt.Fprint(w, " there</h1></body></html>")
}

func main() {
	http.Handle("/", http.HandlerFunc(handle))
    err := http.ListenAndServe(":9002", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}