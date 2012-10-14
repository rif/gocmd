package main

import(
	"log"
	"code.google.com/p/goconf/conf"
)

func main(){
	c, err := conf.ReadConfigFile("something.config")
	if err != nil {
		log.Panic("cannot open config file: ", err)
	}
	log.Print(c.GetString("default", "host")) // return something.com
	log.Print(c.GetInt("default", "port")) // return 443
	log.Print(c.GetBool("default", "active")) // return true
	log.Print(c.GetBool("default", "compression")) // return false
	log.Print(c.GetBool("default", "compression")) // returns false
	log.Print(c.GetBool("service-1", "compression")) // returns true
	log.Print(c.GetBool("service-2", "compression")) // returns GetError
}
