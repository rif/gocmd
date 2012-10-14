package main

import (
	"bufio" 	
 	"os" 	
	"encoding/gob"
 	"fmt"
 	"labix.org/v2/mgo/bson"
)

type User struct {
	ID       bson.ObjectId
	Username string
	Password []byte
	Posts    int
}

func main() {

    in := bufio.NewReader(os.Stdin)    
		dec := gob.NewDecoder(in) 
        u := &User{}        
        dec.Decode(u)
        fmt.Printf("%+v\n", u)
}
