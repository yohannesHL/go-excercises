package diffsquares

// SquareOfSum returns the squares of sums of the first n natural numbers.
func SquareOfSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares returns the sum of square of first n natural numbers.
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference returns the difference between SquareOfSum and SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
