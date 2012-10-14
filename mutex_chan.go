package main

import (
	"log"
	"time"
)

type Command struct {}

func (c *Command) Execute() {
	time.Sleep(5*time.seconds)
}

type ChanMutex struct{
	mu map[string]chan string
}

func NewChanMutex() *ChanMutex {
	return &ChanMutex{mu: make(map[string]chan string}
}

func main() {
	
}
