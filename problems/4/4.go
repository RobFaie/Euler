package main

import "fmt"

func main() {
	fmt.Println("4")

	var p []int

	for i := 500; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			n := i * j
			ns := fmt.Sprintf("%d", n)
			// fmt.Printf("%d %d => %d (%s)\n", i, j, n, ns)

			isPalindrome := true

			for k := 0; k <= (len(ns) / 2); k++ {
				if ns[k] != ns[len(ns)-k-1] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome {
				p = append(p, n)
			}
		}
	}

	max := 0
	for _, v := range p {
		if v > max {
			max = v
		}
	}

	fmt.Printf("%d", max)
}
