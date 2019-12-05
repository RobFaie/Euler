package lib

// GetFibonacci returns fibonacci numbers starting with a and b until reaching limit l
func GetFibonacci(a, b, l int) []int {

	f := []int{a}

	for b <= l {
		a, b = b, a+b
		f = append(f, a)
	}

	return f
}
