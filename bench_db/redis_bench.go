package main

import (
    "fmt"
    "github.com/simonz05/godis"
    "time"
    "log"    
)

func main() {
    r := godis.New("", 10, "")
   	start := time.Now()

    for i:= 0; i < 10e4; i++ {
        r.Set(fmt.Sprintf("a%v", i), i)
    }
    for i:= 0; i < 10e4; i++ {
    	r.Get(fmt.Sprintf("a%v", i))
    }
    for i:= 0; i < 10e4; i++ {
    	r.Del(fmt.Sprintf("a%v", i))
    }
    duration := time.Since(start)
    log.Printf("Elapsed: %v.", duration)
}
