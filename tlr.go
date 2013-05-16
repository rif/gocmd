package main

import (
	"github.com/rif/taller"
	"os"
)

func main() {
	os.Setenv("TALLER_PATH", "/home/rif/Documents/prog/go/src/cmd")
	println(string(taller.Render(taller.NewTemplateFile("template.html"), taller.Context{})))
}
