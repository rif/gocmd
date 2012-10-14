package main

import (
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsRename() {
					log.Println("event:", ev)
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	exit := make(chan bool)
	<-exit
	watcher.Close()
}
