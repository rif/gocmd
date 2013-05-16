package main

import (
	"github.com/fzzy/radix/redis"
	"log"
	"time"
)

func someFunc(c *redis.Client) (err error) {
	reply := c.Cmd("set", "mykey", "myvalue")
	// what is the recommended error
	if reply.Err != nil {
		return reply.Err
	}
	// some code here
	return
}

func main() {
	c, err := redis.DialTimeout("tcp", "192.168.0.17:6379", 5*time.Second)
	if err != nil {
		log.Fatal("Cannot connect to redis: ", err)
	}
	c.Cmd("select", 11)
	defer c.Close()

	err = someFunc(c)
	log.Print("after: ", err, err == nil)
	r := c.Cmd("get", "mykey")
	log.Print(r.Str())
}
