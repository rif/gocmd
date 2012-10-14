package main

import (
	"log"
	"time"
)

func main() {
	d, _ := time.ParseDuration("2s")	
	go func() {
		for {
			t := time.NewTimer(d)
			select {
			case tm := <-t.C:
				log.Printf("Test %v", tm)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	log.Print("End!")
}
