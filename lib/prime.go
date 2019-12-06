package lib

import "math/big"

// IsPrime determines if a number is prime
func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i < n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// GetNextPrime generates the next prime number
func GetNextPrime(x int) int {
	for {
		x++
		if big.NewInt(int64(x)).ProbablyPrime(0) {
			return x
		}
	}
}

// GetPrimeFactorization calculates the prime factorization for a number
func GetPrimeFactorization(n int) []int {
	p := 2

	var f []int

	for p <= n {
		for n%p == 0 {
			f = append(f, p)
			n = n / p
		}
		p = GetNextPrime(p)
	}

	return f
}
