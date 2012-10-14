package main

import (
	"log"
	"log/syslog"
)

func main() {
	logger, err := syslog.NewLogger(syslog.LOG_ALERT, log.LstdFlags)
	if err != nil {
		log.Print("cannot create writer")
	}
	logger.Print("hello syslog world!")
}
