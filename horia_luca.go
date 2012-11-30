package main

import (
	"fmt"
	"time"
)

func main() {
	birthdayRa, _ := time.Parse("Jan 2 2006", "Jul 10 1978")
	birthdayLi, _ := time.Parse("Jan 2 2006", "Oct 31 1980")
	birthdayH, _ := time.Parse("Jan 2 2006", "Aug 18 2008")
	birthdayL, _ := time.Parse("Jan 2 2006", "Dec 3 2010")
	ageRa := time.Since(birthdayRa)
	ageLi := time.Since(birthdayLi)
	ageH := time.Since(birthdayH)
	ageL := time.Since(birthdayL)
	fmt.Printf("Radu is %d days old\n", ageRa/(time.Hour*24))
	fmt.Printf("Liana is %d days old\n", ageLi/(time.Hour*24))
	fmt.Printf("Horia is %d days old\n", ageH/(time.Hour*24))
	fmt.Printf("Luca is %d days old\n", ageL/(time.Hour*24))
}
