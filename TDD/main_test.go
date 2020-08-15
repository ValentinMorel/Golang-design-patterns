package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	a := 5
	b := 6

	expected := 11
	result := sum(a, b)

	if result != expected {
		fmt.Printf("sum function is not working properly; expected : %d, result : %d", expected, result)
	}
}
