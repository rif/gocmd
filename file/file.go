package main

import (
	"os"
	"bufio"
	"fmt"
)


func main() {		
	fin, _ := os.Open("file.go")
	r := bufio.NewReader(fin)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(l))
	}
}

