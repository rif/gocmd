package main

import (
	"fmt"
	"time"
)

func main() {
	format := "2006-1-2 15:04:05 MST"
	now := time.Now()
	y, m, d := now.Date()
	zName, _ := now.Zone()
	l := fmt.Sprintf("%d-%d-%d %s %s", y, m, d, "10:41:00", zName)
	fmt.Println(l)
	p, err := time.Parse(format, l)
	fmt.Println(p)
	fmt.Println(err)
}
