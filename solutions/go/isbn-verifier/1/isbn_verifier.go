package isbnverifier

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i := range len(isbn) {
		val := 0
		if isbn[i] == 'X' && i == 9 {
			val = 10
		} else {
			nVal, err := strconv.Atoi(string(isbn[i]))
			if err != nil {
				return false
			}
			val = nVal
		}
		sum += val * (10 - i)
	}
	return sum%11 == 0
}
