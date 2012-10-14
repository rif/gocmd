package main

import (
	"runtime"
	"fmt"
)

func main(){
	fmt.Println("Number of cpus:", runtime.NumCPU())
}