package main

import (	
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {	
	fmt.Println(rand.Int(rand.Reader, big.NewInt(1000000000)))
}
