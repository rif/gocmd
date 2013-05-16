package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Photo struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Title       string
	Description string
	Image       string
	Active      bool
	User        bson.ObjectId
	UpdatedOn   time.Time
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("photos")
	job := &mgo.MapReduce{
		Map:    "function() { emit(this.user, 1) }",
		Reduce: "function(key, values) { return values.slice(-1)[0] }",
	}
	var result []struct {
		Id    bson.ObjectId "_id"
		Value int
	}
	_, err = c.Find(nil).MapReduce(job, &result)
	if err != nil {
		panic(err)
	}
	var uniqueIds []bson.ObjectId
	for _, item := range result {
		fmt.Println(item)
		uniqueIds = append(uniqueIds, item.Id)
	}
	fmt.Println("Unique:", uniqueIds)
	var photos []*Photo
	c.Find(bson.M{"_id": bson.M{"$in": uniqueIds}}).All(&photos)
	for _, p := range photos {
		fmt.Println("Photo:", p)
	}
}
