package main

func fact(n int) int {
	if  n <= 0 { return 1}
	return n * fact(n-1)
}

func main() {
	for  i:=0; i<1000; i++ {
		fact(i)
	}
}