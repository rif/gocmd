package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	"encoding/json"
)

func printer(m map[string]interface{}, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(indent, k, "=", vv)
		case []interface{}:
			fmt.Println(indent, k, "=")
			for _, u := range vv {
				mm := u.(map[string]interface{})
				for key, val := range mm {
					fmt.Println(key, "=", val)
				}
			}
		case map[string]interface{}:
			level++
			printer(vv, level)
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func cdrHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var f interface{}
	if err := json.Unmarshal(body, &f); err == nil {
		m := f.(map[string]interface{})
		printer(m, 0)
	} else {
		log.Fatalf("Could not unmarshal cdr: %v", err)
	}

}

func main() {
	http.HandleFunc("/cdr", cdrHandler)
	http.ListenAndServe(":8080", nil)
}
