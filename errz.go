package main

import (
	"errors"
	"log"
)

func e() (int, error) {
	return 0, errors.New("test")
}

func beta() (err error) {
	if v, err := e(); err != nil {
		return
	}
	return
}

func main() {
	var err error
	log.Print(err, err != nil)
	log.Print(beta())
}
