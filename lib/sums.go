package lib

// SumOfSquares calculates the sum of squares for a sequence
func SumOfSquares(ns []int) int {
	sum := 0
	for _, v := range ns {
		sum += v * v
	}
	return sum
}

// SquareOfSum calculates the square of the sum of a sequence
func SquareOfSum(ns []int) int {
	sum := 0
	for _, v := range ns {
		sum += v
	}
	return sum * sum
}
