package main

import "fmt"

func main() {
	for i := 1; i < 1000; i++ {
		for j := i + 1; j < 1000; j++ {
			k := 1000 - i - j
			fmt.Printf("i: %d j: %d k:%d\n", i, j, k)
			if (i*i)+(j*j) == (k * k) {
				if i+j+k == 1000 {
					fmt.Printf("%d %d %d => %d", i, j, k, i*j*k)
					return
				}
			}
		}
	}
}
