package main

import (
	"code.google.com/p/gocask"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	g, _ := gocask.NewGocask("testkv")
	defer g.Close()
	start := time.Now()

	for i := 0; i < 10e4; i++ {
		g.Put(fmt.Sprintf("a%v", i), []byte(strconv.Itoa(i)))
	}
	for i := 0; i < 10e4; i++ {
		g.Get(fmt.Sprintf("a%v", i))
	}
	for i := 0; i < 10e4; i++ {
		//r.Del(fmt.Sprintf("a%v", i))
	}
	duration := time.Since(start)
	log.Printf("Elapsed: %v.", duration)
}
