package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	result := sum(a, b)

	fmt.Printf("sum of %d and %d is %d \n", a, b, result)
}
