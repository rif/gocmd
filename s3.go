package main

import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
)

func main() {
	connection := s3.New(aws.Auth{"*************", "****************"}, aws.USEast)

	log.Print(connection.Bucket("lov3ly").URL("/"))
}
