package main

import (
	"fmt"

	"../../lib"
)

func main() {

	s := 0

	fib := lib.GetFibonacci(1, 2, 4000000)

	fmt.Printf("Fib: %s\n", fib)

	for _, v := range fib {
		if v%2 == 0 {
			s += v
		}
	}

	fmt.Printf("Sum: %d\n", s)
}
