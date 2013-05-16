package main

import (
	"crypto/aes"
	. "fmt"
	"os"
)

func main() {
	msgbuf := []byte("This is long message text. len32 nsandas askh skljhklj sdfhlkh ds")
	// some key, 16 Byte long
	key := []byte("mama are mere!!!")

	// create the new cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		Println("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}
	out := make([]byte, len(msgbuf))

	c.Encrypt(msgbuf, out)
	Println(">> ", out)

	// now we decrypt our encrypted text
	plain := make([]byte, len(out))
	c.Decrypt(out, plain) // decrypt the first half

	Println("msg: ", string(plain))
}
