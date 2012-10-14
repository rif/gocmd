package main

import (
	"fmt"
    "github.com/fsouza/gokabinet/kc"
    "time"
    "log"
)

func main() {
    db, _ := kc.Open("cache.kch", kc.WRITE)
    defer db.Close()
	start := time.Now()
    for i:= 0; i < 10e4; i++ {
        db.SetInt(fmt.Sprintf("a%v", i), i)
    }
    for i:= 0; i < 10e4; i++ {
    	db.Get(fmt.Sprintf("a%v", i))
    }
    for i:= 0; i < 10e4; i++ {
    	db.Remove(fmt.Sprintf("a%v", i))
    }
    duration := time.Since(start)
    log.Printf("Elapsed: %v.", duration)
}
