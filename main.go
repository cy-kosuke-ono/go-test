package main

import (
	"fmt"
	"os"
)

func add(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func prod(a int, b int) int {
	return a * b
}

func main() {
	fmt.Printf("add: %v\n", add(1, 4))
	fmt.Printf("minus: %v\n", minus(6, 4))
	os.Exit(0)
}
