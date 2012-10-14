package main

import (
	"github.com/cgrates/cgrates/inotify"
	"log"	
)

func main() {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch(".")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			if ev.Mask&inotify.IN_MODIFY != 0 {
				log.Println("event:", ev)
				log.Println("file name:", ev.Name)
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
