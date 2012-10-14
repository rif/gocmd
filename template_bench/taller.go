package main

import (
	"github.com/rif/taller"
	"io"
	"log"
	"net/http"
	"os"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	p, _ := os.Getwd()
	os.Setenv("TALLER_PATH", p+"/taller")
	template := taller.NewTemplateFile("index.html")
	context := taller.Context{"c": os.Getenv("TALLER_PATH")}

	content := taller.Render(template, context)
	io.WriteString(w, string(content))
}

func main() {
	http.HandleFunc("/", mainPage)
	log.Print("serving at 8000...")
	http.ListenAndServe(":8000", nil)
}
