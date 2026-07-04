package sumofmultiples

import (
	"slices"
)

func SumMultiples(limit int, divisors ...int) int {
	var multiples []int
	for _, div := range divisors {
		if div == 0 {
			continue
		}
		for i := 1; i < limit; i++ {
			if i%div == 0 {
				if !slices.Contains(multiples, i) {
					multiples = append(multiples, i)
				}
			}
		}
	}
	sum := 0
	for _, x := range multiples {
		sum += x
	}
	return sum
}
