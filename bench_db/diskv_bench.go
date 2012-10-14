package main

import (
	"fmt"
	"github.com/peterbourgon/diskv"
	"log"
	"strconv"
	"time"
)

func main() {
	// Simple transform function to put all of the data files into the root directory.
	flatTransform := func(s string) []string { return []string{""} }

	// Initialize a new diskv store, rooted at "my-store-dir", with a 1MB cache.
	s := diskv.NewStore("my-store-dir", flatTransform, 1024*1024)

	start := time.Now()

	for i := 0; i < 100; i++ {
		s.Write(fmt.Sprintf("a%v", i), []byte(strconv.Itoa(i)))
	}
	for i := 0; i < 100; i++ {
		s.Read(fmt.Sprintf("a%v", i))
	}
	for i := 0; i < 100; i++ {
		s.Erase(fmt.Sprintf("a%v", i))
	}
	duration := time.Since(start)
	log.Printf("Elapsed: %v.", duration)
}
