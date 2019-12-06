package main

import (
	"fmt"

	"../../lib"
)

func main() {
	fmt.Println("Problem 10")

	p := 2
	sum := 0
	for p < 2000000 {
		sum += p
		p = lib.GetNextPrime(p)
	}

	fmt.Printf("Sum: %d", sum)
}
