package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	rune, _, err := in.ReadRune()
	if err != nil {
		log.Print("my error: ", err)
	}
	fmt.Println("*", string(rune))
}
