package main

import (
        "fmt"
        "launchpad.net/mgo/bson"
)

type Person struct {
    Name string
    Phone string ",omitempty"
}

func main() {
        data, err := bson.Marshal(&Person{Name:"Bob"})
        if err != nil {
                panic(err)
        }
	p := new(Person)
	bson.Unmarshal(data, p)
        fmt.Printf("%v\n", p.Name)
}
