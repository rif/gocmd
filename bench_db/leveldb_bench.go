package main

import (
	"fmt"
	"github.com/jmhodges/levigo"
	"log"
	"strconv"
	"time"
)

func main() {
	dbname := "leveldb"
	opts := levigo.NewOptions()
	opts.SetCache(levigo.NewLRUCache(3 << 20))
	opts.SetCreateIfMissing(true)
	_ = levigo.DestroyDatabase(dbname, opts)
	db, _ := levigo.Open(dbname, opts)

	wo := levigo.NewWriteOptions()
	ro := levigo.NewReadOptions()

	start := time.Now()

	for i := 0; i < 10e4; i++ {
		db.Put(wo, []byte(fmt.Sprintf("a%v", i)), []byte(strconv.Itoa(i)))
	}
	for i := 0; i < 10e4; i++ {
		db.Get(ro, []byte(fmt.Sprintf("a%v", i)))
	}
	for i := 0; i < 10e4; i++ {
		db.Delete(wo, []byte(fmt.Sprintf("a%v", i)))
	}
	duration := time.Since(start)
	log.Printf("Elapsed: %v.", duration)
}
