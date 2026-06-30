package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
		return hammingDistance, nil
	}
	return hammingDistance, errors.New("strings must be of equal length")
}
