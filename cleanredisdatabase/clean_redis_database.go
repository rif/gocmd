package main

import (
    "fmt"
    "os"
    "log"
    "strconv"
    "github.com/simonz05/godis"    
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n\tclean_redis_database <database_nb>")
		log.Fatal("Bye!")
	}	
	db_nb, _ := strconv.Atoi(os.Args[1])
    r := godis.New("",db_nb , "")
    keys,_ := r.Keys("*")
    for _,key := range(keys) {
    	r.Del(key)
    }
    fmt.Println("done!")
}
