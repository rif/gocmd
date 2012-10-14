package main

import (
		"log"
		"github.com/simonz05/godis"
		"time"		
)

func main(){
	r := godis.New("", 10, "")
	r.Rpush("lista", 1)
	r.Rpush("lista", 2)
	r.Rpush("lista", 3)
	result,_ := r.Lrange("lista",-100,100)
	r.Hmset("hash", map[string]interface{}{"month":1, "now":time.Now()})
	log.Print(result.IntArray())
	result1,_ := r.Hgetall("hash")
	m:=result1.StringMap() 
	log.Print(time.Parse(m["now"],m["now"]))
}
