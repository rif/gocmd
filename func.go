package main

import (
	"log"
)

type actionTypeFunc func(balanceId string, units float64)

func logAction(balanceId string, units float64) {
	log.Printf("Balance: %s units %v", balanceId, units)
}

var (
	actionTypeMap = map[string]actionTypeFunc{
		"test": logAction,
	}
)

func main() {
	actionTypeMap["test"]("hello", 1.1)
}
