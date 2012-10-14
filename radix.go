package main

import (
	"github.com/fzzbt/radix/redis"
	"log"
)

func someFunc(c *redis.Client)(err error){
	reply := c.Set("mykey", "myvalue")
	// what is the recommended error 
	if reply.Err != nil {
		return reply.Err
	}
	// some code here
	return
}

func main() {
	conf := redis.DefaultConfig()
	c := redis.NewClient(conf)
	defer c.Close()

	err :=  someFunc(c)
	log.Print("after: ", err, err == nil)
}
