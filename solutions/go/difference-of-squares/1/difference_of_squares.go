package differenceofsquares

import "math"

func SquareOfSum(n int) int {
	sum := 0
	for i := range n + 1 {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := range n + 1 {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	diff := SumOfSquares(n) - SquareOfSum(n)
	return int(math.Abs(float64(diff)))
}
