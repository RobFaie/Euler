package main

import (
	"fmt"

	"../../lib"
)

func main() {
	fmt.Println("Problem 6.")

	var r []int
	for i := 1; i <= 100; i++ {
		r = append(r, i)
	}

	a := lib.SumOfSquares(r)
	b := lib.SquareOfSum(r)

	fmt.Printf("%d %d => %d", a, b, b-a)
}
