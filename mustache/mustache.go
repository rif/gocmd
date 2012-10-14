package main

import (
	"github.com/hoisie/mustache"
	"log"
)

func main(){
	out := mustache.RenderFileInLayout("template.html.mustache", "layout.html.mustache", nil)
	log.Print(out)
	out = mustache.RenderFileInLayout("t1.html.mustache", "l1.html.mustache", map[string]string{"title": "Title of the page",
  "content": "This is some content", "copyright": "&copy;2012 Foo"})
	log.Print(out)
}
