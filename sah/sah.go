package main

import (
	"fmt"
	
	"crypto/sha1"
)

func main() {
	a := sha1.New()
	//io.Copy(a, strings.NewReader("mama"))
	fmt.Printf("sha1: %x\n", a.Sum([]byte("mama")))
}