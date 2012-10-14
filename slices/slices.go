package main

import (
	"fmt"
)

type Slicer struct{
	a []string
}

func main(){
	s := Slicer{}
	s.a = append(s.a, "Mama")
	s.a = append(s.a, "Tata")	
	fmt.Println(s.a);
}