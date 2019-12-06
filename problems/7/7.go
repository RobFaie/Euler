package main

import (
	"fmt"

	"../../lib"
)

func main() {
	fmt.Println("Problem 7.")

	p := 2

	for i := 1; i <= 10001; i++ {
		fmt.Printf("Prime #%d: %d\n", i, p)
		p = lib.GetNextPrime(p)
	}

	// fmt.Printf("Prime #10 001: %d", p)
}

// 104729
