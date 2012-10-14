package main

import (
	"fmt"
	"time"
	"log"
)

func main(){
	m := make(map[string]int)
	start := time.Now() 
    for i:= 0; i < 10e4; i++ {
        m[fmt.Sprintf("a%v", i)] = i
    }
    for i:= 0; i < 10e4; i++ {
    	v:= m[fmt.Sprintf("a%v", i)]
    	v++
    }
    for i:= 0; i < 10e4; i++ {
    	delete(m, fmt.Sprintf("a%v", i))
    }
    duration := time.Since(start)
    log.Printf("Elapsed: %v.", duration)
}

