package main

import (
	"log"
	"math/rand"
)

func pwdgen(length int) string {
	alphabet := "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	password := make([]byte, length)
	for i := 0; i < len(password); i++ {
		password[i] = alphabet[rand.Int()%len(alphabet)]
	}
	return string(password)
}

func main() {
	for i := 0; i < 10; i++ {
		log.Print(pwdgen(8))
	}
}
