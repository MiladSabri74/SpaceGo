package armstrongnumbers

import "math"

func CountDigit(n int) int {
	if n == 0 {
		return 1
	}
	r := 0
	for n >= 1 {
		n = n / 10
		r++
	}
	return r
}
func IsNumber(n int) bool {
	x := CountDigit(n)
	value := n
	armstrong := 0

	for n >= 1 {
		armstrong += int(math.Pow(float64(n%10), float64(x)))
		n = n / 10
	}
	if armstrong == value {
		return true
	}
	return false
}
