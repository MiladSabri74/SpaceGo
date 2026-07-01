package raindrops

import (
  "strconv"
)
func Convert(number int) string {
	result := ""
	isDivisible := false
	if (number % 3 == 0) {
		result +="Pling"
		isDivisible = true
	}
	if (number % 5 == 0) {
		result +="Plang"
		isDivisible = true
	}
	if (number % 7 == 0) {
		result +="Plong"
		isDivisible = true
	}
	if (!isDivisible) {
		return strconv.Itoa(number)
	}
	return result
}
