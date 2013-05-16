package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type Cost struct {
	Call  string
	Price float64
}

type Log struct {
	Id   string `_id,omitempty`
	Cost *Cost
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("log")
	c.Upsert(bson.M{"_id": "test1"}, &Log{"test1", &Cost{"mama", 10}})

	result := Log{}
	err = c.Find(bson.M{"_id": "test1"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Result:", result)
}
