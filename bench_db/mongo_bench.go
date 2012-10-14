package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type KV struct {
	Key   string
	Value int
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("bench")
	start := time.Now()

	for i := 0; i < 10e4; i++ {
		c.Insert(&KV{fmt.Sprintf("a%v", i), i})
	}
	result := KV{}
	for i := 0; i < 10e4; i++ {
		c.Find(bson.M{"key": fmt.Sprintf("a%v", i)}).One(&result)
	}
	for i := 0; i < 10e4; i++ {
		c.Remove(bson.M{"key": fmt.Sprintf("a%v", i)})
	}
	duration := time.Since(start)
	log.Printf("Elapsed: %v.", duration)
}
