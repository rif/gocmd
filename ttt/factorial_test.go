package main

import (
	"testing"
)

func Test0(t *testing.T) {
	if factorial(10) != 3628800 {
		t.Error("Buba in factorial.")
	}
}