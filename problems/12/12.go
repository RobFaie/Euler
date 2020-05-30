package main

import "fmt"

func main() {
	fmt.Println("Problem 12.")

	n := 0
	for i := 1; true; i++ {
		n += i
		var factors []int

		for j := 1; j <= n; j++ {
			if n%j == 0 {
				factors = append(factors, j)
			}
		}

		fmt.Printf("Triangular number: %d => %d\n", n, len(factors))
		// fmt.Printf("Factors: %s\n", factors)

		if len(factors) > 500 {
			fmt.Printf("Num: (%d) %d", i, n)
			return
		}
	}
}
