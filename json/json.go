package main

import (
	"log"
	"encoding/json"
)

type Destination struct {
	Id       string
	Prefixes []string
}

func main(){
	d := Destination{"mama", []string{"0723045326", "0723"}}
	
}
