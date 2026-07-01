package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	newId, _ := strconv.Atoi(id)
	if newId == 0 && len(id) == 1 {
		return false
	}
	sum := 0
	index := 1
	for i := len(id) - 1; i >= 0; i-- {
		if index%2 == 0 {
			val, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return false
			}
			val = val * 2
			if val > 9 {
				val -= 9
			}
			sum += val
		} else {
			val, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return false
			}
			sum += val
		}
		index++
	}
	if sum%10 == 0 {
		return true
	}

	return false
}
