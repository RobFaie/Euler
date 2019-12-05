package main

import (
	"fmt"

	"../../lib"
)

func main() {

	n := 600851475143
	p := 2

	fmt.Printf("Prime factors:")

	for p <= n {
		for n%p == 0 {
			fmt.Printf(" %d", p)
			n = n / p
		}
		p = lib.GetNextPrime(p)
	}

}
