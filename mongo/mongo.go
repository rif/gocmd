package main

import (
        "fmt"
        "launchpad.net/mgo/bson"
        "launchpad.net/mgo"
        "time"
)

type Interval struct {
	Now time.Time
}

type Person struct {
        Name string
        Phone string
        Other map[string]Interval
}

func main() {
        session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"ole", "+55 53 8116 9639", map[string]Interval{"1": Interval{time.Now()}}},
	               &Person{"Cla", "+55 53 8402 8510", map[string]Interval{"1": Interval{time.Now()}}})
        if err != nil {
                panic(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "ole"}).One(&result)
        if err != nil {
                panic(err)
        }

        fmt.Println("Phone:", result.Other["1"].Now.Month())
}
