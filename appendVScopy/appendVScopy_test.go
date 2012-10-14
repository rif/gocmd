package main

import ("testing")

func BenchmarkAppend(b *testing.B){
	x:=[]int{1,2,3,4,5,6,7,8,9,0}	
	for i := 0; i < b.N; i++ {
		var y []int	
		y = append(y,x...)		
	}	
}

func BenchmarkCopy(b *testing.B){
	x:=[]int{1,2,3,4,5,6,7,8,9,0}	
	for i := 0; i < b.N; i++ {
		y := make([]int, len(x))	
		copy(y,x)		
	}	
}
