package main

import "fmt"

func main() {
	fmt.Println("Problem 5.")

	for i := 1; true; i++ {
		valid := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Printf(">%d", i)
			return
		}
	}
}
