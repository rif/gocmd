package main

import (
	"fmt"
	"github.com/kklis/gomemcache"
	"time"
	"log"
)

func main() {
    mc := Connect("127.0.0.1", 11211)
   	start := time.Now()
    for i:= 0; i < 10e4; i++ {
        mc.Set(fmt.Sprintf("a%v", i), i)
    }
    for i:= 0; i < 10e4; i++ {
    	mc.Get(fmt.Sprintf("a%v", i))
    }
    for i:= 0; i < 10e4; i++ {
    	mc.Delete(fmt.Sprintf("a%v", i))
    }
    duration := time.Since(start)
    log.Printf("Elapsed: %v.", duration)
}
