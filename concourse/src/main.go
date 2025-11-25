package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		sum := add(a, b)
		fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)
	} else {
		fmt.Println("Please provide two arguments.")
		os.Exit(1)
	}

}

func add(a, b int) int {
	return a + b
}
