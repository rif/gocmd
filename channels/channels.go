package main

import (
	"log"
	"time"
)

func DebitLoop(stop chan byte) {
	for {
		select {
		case <-stop:
			log.Print("Put credit back!")
			return
		default:
		}
		log.Print("Tic")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan byte)
	go DebitLoop(ch)
	time.Sleep(5 * time.Second)
	ch <- 1
	log.Print("Done!")
}
